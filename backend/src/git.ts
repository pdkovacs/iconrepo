import * as path from "path";
import { spawn, exec, SpawnOptions } from "child_process";
import { Observable, Observer } from "rxjs";
import { mkdirMaybe, appendFile } from "./utils/rx";
import {
    SerializableJobImpl,
    JobDoneCallback,
    JobResult,
    create as createSerializer
} from "./utils/serializer";
import logger, { ContextAbleLogger } from "./utils/logger";
import { CreateIconInfo } from "./icon";
import { commandExecutor } from "./utils/command-executor";

type GitOperationCallback = (error: Error) => void;
type GitCommandExecutor = (spawnArgs: string[]) => Observable<string>;

export const GIT_COMMIT_FAIL_INTRUSIVE_TEST = "GIT_COMMIT_FAIL_INTRUSIVE_TEST";

export const createGitCommandExecutor: (
    iconRepository: string
) => GitCommandExecutor
= iconRepository => spawnArgs => {
    const ctxLogger = logger.createChild(`executeGitCommand ${spawnArgs} in ${iconRepository}`);
    return commandExecutor(ctxLogger, "git", spawnArgs, { cwd: iconRepository });
};

const enqueueJob = createSerializer("GIT");

const getFileName: (inconFileInfo: CreateIconInfo) => string
    = inconFileInfo => `${inconFileInfo.name}@${inconFileInfo.size}.${inconFileInfo.format}`;

/*
 * @return an Observable for the path to the icon file relative to the local GIT repository's root.
 */
const createIconFile: (inconFileInfo: CreateIconInfo, iconRepository: string) => Observable<string>
= (inconFileInfo, iconRepository) =>
    mkdirMaybe(path.join(iconRepository, inconFileInfo.format))
    .flatMap(pathToFormatDir =>
        mkdirMaybe(path.join(pathToFormatDir, inconFileInfo.size))
        .flatMap(pathToSizeDir =>
            appendFile(
                path.join(
                    pathToSizeDir,
                    getFileName(inconFileInfo)
                ),
                inconFileInfo.content,
                { flag: "w"})
            ));

const addToIndex = (pathInRepo: string) => ["add", pathInRepo];

const commitCommand: () => string = () =>
    process.env.GIT_COMMIT_FAIL_INTRUSIVE_TEST === "true"
        ? "procyon lotor"
        : "commit";

const commit: (pathInRepo: string, userName: string) => string[]
= (pathInRepo, userName) => [
    commitCommand(),
    "-m", `"${pathInRepo}" icon added by ${userName}`,
    `--author=CXN Icon Repository Server <${userName}>`
];

const rollback: () => string[][] = () => [
    ["reset", "--hard", "HEAD"],
    ["clean", "-qfdx"]
];

const createAddIconFileJob: (
    inconFileInfo: CreateIconInfo,
    userName: string,
    gitExec: GitCommandExecutor,
    iconRepository: string,
    callback: GitOperationCallback
) => SerializableJobImpl
= (inconFileInfo, userName, gitExec, iconRepository, callback) => {
    const ctxLogger = logger.createChild("add icon file");
    ctxLogger.debug("BEGIN");
    return (done: JobDoneCallback) =>
            createIconFile(inconFileInfo, iconRepository)
            .flatMap(pathInRepo =>
                gitExec(addToIndex(pathInRepo))
                .flatMap(() =>
                    gitExec(commit(pathInRepo, userName)))
                .mapTo(pathInRepo))
            .subscribe(
                pathInRepo => {
                    callback(void 0);
                    done();
                },
                error => {
                    ctxLogger.error(`Adding file failed with ${error}`);
                    gitExec(rollback()[0])
                    .flatMap(() => gitExec(rollback()[1]))
                    .catch(errorInRollback => {
                        ctxLogger.error(errorInRollback);
                        return "dummy return value";
                    })
                    .subscribe(
                        void 0,
                        e => {
                            callback(e);
                            done();
                        },
                        () => {
                            callback(error);
                            done();
                        }
                    );
                },
                void 0
            );
};

type AddIconFile = (
    inconFileInfo: CreateIconInfo,
    userName: string
) => Observable<void>;

export interface GitAccessFunctions {
    readonly getRepoLocation: () => string;
    readonly addIconFile: AddIconFile;
}

type GitAFsProvider = (localIconRepositoryLocation: string) => GitAccessFunctions;

const gitAccessFunctionsProvider: GitAFsProvider = localIconRepositoryLocation => ({
    getRepoLocation: () => localIconRepositoryLocation,

    addIconFile: (inconFileInfo, userName) => Observable.create((observer: Observer<string>) =>
    enqueueJob(
            createAddIconFileJob(
                inconFileInfo,
                userName,
                createGitCommandExecutor(localIconRepositoryLocation),
                localIconRepositoryLocation,
                (error: Error) => {
                    if (error) {
                        observer.error(error);
                    } else {
                        observer.next(void 0);
                        observer.complete();
                    }
                }
            )
        )
    )
});

export default gitAccessFunctionsProvider;

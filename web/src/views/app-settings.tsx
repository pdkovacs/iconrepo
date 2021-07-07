import * as React from "react";
import { VersionInfo } from "../services/config";
import "./app-settings.scss";

import { Dialog, Classes, Button } from "@blueprintjs/core";

const dialogOptions = {
    autoFocus: true,
    canEscapeKeyClose: true,
    canOutsideClickClose: true,
    enforceFocus: true,
    hasBackdrop: true,
    title: "About \"Icons\"",
    usePortal: true
};

export class AppSettgins extends React.Component<{versionInfo: VersionInfo}, {isOpen: boolean}> {

    constructor(props: {versionInfo: VersionInfo}) {
        super(props);
        this.state = { isOpen: false };
    }

    public render() {
        return <div>
            <h1 className="app-settings">
                <a onClick={this.handleOpen}>Icons</a>
            </h1>
            <Dialog
                {...dialogOptions}
                isOpen={this.state.isOpen}
                onClose={this.handleClose}>
                <div className={Classes.DIALOG_BODY}>
                    <div className="about-dialog">
                        <table>
                            <tbody>
                                <tr>
                                    <td>Version:</td><td>{this.props.versionInfo.version}</td>
                                </tr>
                                <tr>
                                    <td>Commit ID:</td><td>{this.props.versionInfo.commit}</td>
                                </tr>
                            </tbody>
                        </table>
                    </div>
                </div>
                <div className={Classes.DIALOG_FOOTER}>
                    <div className={Classes.DIALOG_FOOTER_ACTIONS}>
                        <Button onClick={this.handleClose}>Close</Button>
                    </div>
                </div>
            </Dialog>
        </div>;
    }

    private handleOpen = () => this.setState({ isOpen: true });
    private handleClose = () => this.setState({ isOpen: false });
}

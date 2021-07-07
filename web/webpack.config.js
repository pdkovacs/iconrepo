const path = require('path');
require('webpack');

const React = require('react');

const HTMLWebpackPlugin = require('html-webpack-plugin');

// const DEVELOPMENT = process.env.NODE_ENV === 'development';
const TEST = process.env.NODE_ENV === 'test';
const PRODUCTION = process.env.NODE_ENV === 'production';

var nodeExternals = require('webpack-node-externals');

const mode = PRODUCTION
    ? 'production'
    : 'development';

const entry = TEST
    ? [ './test/specRoot.js' ]
    : [ // PRODUCTION || DEVELOPMENT
        './src/index.tsx'
    ];

const publicPath = PRODUCTION
    ? './'
    : '/'

const plugins = PRODUCTION
    ? [
            new HTMLWebpackPlugin({
                template: 'index-prod-template.html',
                templateParameters: {
                    'reactVersion': React.version
                },
                hash: true,
                cache: false
			})
    ]
    : TEST
    ? []
    : [ // DEVELOPMENT
            new HTMLWebpackPlugin({
                template: 'index-dev-template.html',
                templateParameters: {
                    'reactVersion': React.version
                },
                hash: true,
                cache: false
			})
    ]

const cssIdentifier = '[path][name]---[local]';

const cssLoader = [
    'style-loader',
    'css-loader?localIdentName=' + cssIdentifier
];

const externals = TEST ? [
    nodeExternals({
        whitelist: ['normalize.css', 'whatwg-fetch']
    })
] : {
    'react': 'React',
    'react-dom': 'ReactDOM'
};

const typeScriptConfigFileName = TEST
    ? path.resolve(__dirname, 'tsconfig-test.json')
    : path.resolve(__dirname, 'tsconfig.json');

module.exports = {
    mode,
    devtool: 'source-map',
    entry: entry,
    plugins: plugins,
    resolve: {
        alias: {
            src: path.resolve(__dirname, 'src')
        },
        // Add '.ts' and '.tsx' as resolvable extensions.
        extensions: ['.ts', '.tsx', '.js', '.json']
    },
    module: {
        rules: [
            // All files with a '.ts' or '.tsx' extension will be handled by 'awesome-typescript-loader'.
            {
                test: /\.tsx?$/,
                loader: 'ts-loader',
                options: {
                    configFile: typeScriptConfigFileName
                }
            },
            // All output '.js' files will have any sourcemaps re-processed by 'source-map-loader'.
            {
                enforce: 'pre',
                test: /\.js$/,
                loader: 'source-map-loader'
            },
            {
                test: /normalize\.css$/,
                loaders: cssLoader
            },
            {
                test: /\.(woff|woff2)$/,
                use: {
                  loader: 'url-loader',
                  options: {
                    name: 'fonts/[hash].[ext]',
                    limit: 5000,
                    mimetype: 'application/font-woff'
                  }
                }
            },
            {
                test: /\.(ttf|eot|svg)$/,
                use: {
                  loader: 'file-loader',
                  options: {
                    name: 'fonts/[hash].[ext]'
                  }
                }
            },
            {
                test: /\.scss$/,
                loaders: ["style-loader", "css-loader", "sass-loader"],
                exclude: /node_modules/
            },
            {
                test: /\.svg$/,
                use: {
                    loader: 'svg-url-loader',
                    options: {limit: 1, noquotes: true}
                }
            }
        ]
    },
    externals: externals,
    output: {
        path: path.join(__dirname, 'dist'),
        publicPath,
        filename: 'bundle.js'
    }
};

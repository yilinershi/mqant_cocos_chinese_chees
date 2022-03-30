npm install -D webpack@4.41.5 webpack-cli@3.3.10 webpack-dev-server@3.10.2

# 使用webpack打包
* 下载依赖：
```
npm install -D 
npm install -D webpack@4.41.5 
npm install -D webpack-cli@3.3.10
npm install -D webpack-dev-server@3.10.2
npm install -D clean-webpack-plugin@3.0.0
npm install -D html-webpack-plugin@4.5.0
npm install -D ts-loader@8.0.11
npm install -D cross-env@7.0.2
```
* 生成package.json:`npm init -y`
* 生成tsconfig.json:`tsc --init`
* 在build目录下新建webpack.config.js文件
``` js
const { CleanWebpackPlugin } = require('clean-webpack-plugin')
const HtmlWebpackPlugin = require('html-webpack-plugin')
const path = require('path')

const isProd = process.env.NODE_ENV === 'production' // 是否生产环境

function resolve(dir) {
    return path.resolve(__dirname, '..', dir)
}

module.exports = {
    mode: isProd ? 'production' : 'development',
    entry: {
        app: './src/main.ts'
    },

    output: {
        path: resolve('dist'),
        filename: '[name].[contenthash:8].js'
    },

    module: {
        rules: [
            {
                test: /\.tsx?$/,
                use: 'ts-loader',
                include: [resolve('src')]
            }
        ]
    },

    plugins: [
        new CleanWebpackPlugin({
        }),

        new HtmlWebpackPlugin({
            template: './public/index.html'
        })
    ],

    resolve: {
        extensions: ['.ts', '.tsx', '.js']
    },

    devtool: isProd ? 'cheap-module-source-map' : 'cheap-module-eval-source-map',

    devServer: {
        host: 'localhost', // 主机名
        stats: 'errors-only', // 打包日志输出输出错误信息
        port: 8081,
        open: true
    },
}

```
* 修改package.json中的script命令
```js
"scripts": {
    "dev": "cross-env NODE_ENV=development webpack-dev-server --config build/webpack.config.js",
    "build": "cross-env NODE_ENV=production webpack --config build/webpack.config.js"
  },
``` 

# 注意
* 如果安装的webpack依赖包时使用最新的，可能会不兼容
* 使用`npm init -y`初始化package.json时，如果项目名中有中文，会导致初始化不成功

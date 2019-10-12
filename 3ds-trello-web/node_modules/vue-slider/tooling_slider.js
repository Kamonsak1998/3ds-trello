import webpack from 'webpack'
export default function (config, options) {
  config.devtool = 'source-map'
  config.entry =  {
    'vue-slider': './index.js'
  }
  config.output = {
    filename: './dist/[name].js',
    library: 'VueSlider',
    libraryTarget: 'umd'
  }
  config.plugins = [
    new webpack.DefinePlugin({
      'process.env': {
        NODE_ENV: '"production"'
      }
    }),
    new webpack.optimize.UglifyJsPlugin({
      compress: {
        warnings: false
      }
    })
  ]
  return config
}
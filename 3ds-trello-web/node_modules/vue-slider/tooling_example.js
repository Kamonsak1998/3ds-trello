export default function (config, options) {
  config.entry = ['./example/index']
  config.module.preLoaders = [
    {
      test: /\.vue$/,
      loader: 'eslint',
      exclude: /node_modules/
    },
    {
      test: /\.js$/,
      loader: 'eslint',
      exclude: /node_modules/
    }
  ]
  config.eslint = {
    formatter: require('eslint-friendly-formatter')
  }
  return config
}
import path from 'path';

export default {
  entry: './src/app',
  output: {
    path: path.resolve(__dirname, 'dist'),
    filename: 'bundle.js',
  },
  module: {
    rules: [{
      test: /\.js$/,
      include: [path.resolve(__dirname, 'src')],
      loader: 'babel-loader',
    }, {
      test: /\.vue$/,
      loader: 'vue-loader',
    }],
  },
};

module.exports = {
  root: true,
  parserOptions: {
    project: './tsconfig.json'
  },
  parser: '@typescript-eslint/parser',
  plugins: [
    '@typescript-eslint',
  ],
  extends: [
    'eslint:recommended',
    'plugin:@typescript-eslint/recommended',
    'airbnb',
    'airbnb-typescript',
  ],
  env: {
    node: true,
    jest: true,
  },
  rules: {
    "react/function-component-definition": [2, { "namedComponents": "arrow-function" }]
  },
  ignorePatterns: ['.eslintrc.js'],
};

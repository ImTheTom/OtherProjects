module.exports = function moduleExport(api) {
  api.cache(true);
  return {
    presets: ['babel-preset-expo'],
  };
};

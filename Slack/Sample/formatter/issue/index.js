const { lookupType } = require('./common');

module.exports = (reports, jpdOrigin, policyOrWatchName) => {
  const mapper = lookupType(__dirname, reports[0].type);
  if (mapper) {
    const results = mapper(reports, jpdOrigin, policyOrWatchName).map((o) => ({
      ...o,
      type: reports[0].type,
    }));

    return results;
  }
};

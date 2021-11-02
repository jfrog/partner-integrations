const { existsSync } = require('fs');
const path = require('path');

const REPO_TYPES = {
  alpine: 'alpine://branch:package:version',
  composer: () => 'composer://package:version',
  deb: () => 'deb://vendor:dist:package:version',
  docker: (report) => {
    const [name, version] = report.artifact.display_name.split(':');
    return {
      name: `docker://${name}/name:tag`,
      version,
      path: `${report.artifact.path.replace('default/', '').split('/').slice(0, -2).join('/')}/`,
    };
  },
  generic: () => 'generic://sha256:<Checksum>/name',
  golang: () => 'go://package:version',
  maven: (report) => `gav://${report.display_name}`,
  npm: () => 'npm://package:version',
  nuget: () => 'nuget://module:version',
  python: () => 'pip://package:version',
  rpm: () => 'rpm://os-version:package:version',
  unknown: (report) => `${report.pkg_type}://${report.display_name}`,
};

const findComponent = (type) => REPO_TYPES[type.toLowerCase()] || REPO_TYPES.unknown;

const fields = (o) => Object.entries(o).map(([k, v]) => ({
  type: 'mrkdwn',
  text: `*${k}*: ${v}`,
  verbatim: true,
}));

const lookupType = (dir, type) => {
  const script = `${path.join(dir, type)}.js`;
  return existsSync(script) ? require(script) : undefined;
};

const longFields = (o) => fields(o).map((f) => ({
  type: 'section',
  text: f,
}));

const lookupFormatMapper = (format) => {
  const formatterPath = path.join(__dirname, format);
  const formatter = existsSync(formatterPath);
  if (!formatter) {
    console.error(`No XR message formatter found ${format}, formatter path: ${formatterPath}`);
  }
  return formatter ? require(path.join(__dirname, format)) : undefined;
};

module.exports = {
  longFields,
  fields,
  findComponent,
  lookupType,
  lookupFormatMapper,
};

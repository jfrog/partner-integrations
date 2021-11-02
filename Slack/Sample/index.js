const formatter = require('./formatter');

const main = (payload, formatType) => {
  const event = {
    violation: payload,
    jpdOrigin: 'https://tempurl.org',
  };

  const results = formatter(event, formatType, event.violation.watch_name);

  const slackMessages = results.map(({ blocks }) => ({ blocks }));

  console.log(JSON.stringify(slackMessages[0], null, 4));
};

if (process.argv.length <= 2) {
  console.log('Usage: node index.js <payload json file path>');
  process.exit(1);
}

const payload = require(`./${process.argv[2]}`);
const formatType = process.argv.length == 4 ? process.argv[3] : 'summary';

main(payload, formatType);

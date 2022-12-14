const fs = require("fs");
const path = require("path");
const spawnSync = require("child_process").spawnSync;

const packages = fs
  .readdirSync(__dirname)
  .filter((v) => v.startsWith("fslint-"))
  .concat(["fslint"]);

for (const pkg of packages) {
  const cwd = path.join(__dirname, pkg);

  spawnSync("npm", ["publish", "--access-public"], {
    cwd: cwd,
    shell: true,
    stdio: "inherit",
    env: {
      NODE_AUTH_TOKEN: process.env.NODE_AUTH_TOKEN,
    },
  });
}

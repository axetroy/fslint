const spawn = require("child_process").spawn;
const os = require("os");
const path = require("path");

const platform = os.platform();
const arch = os.arch();

const ERR_NOT_SUPPORT = new Error("fslint does not support your platform");

const platformMap = {
  win32: {
    ia32: "fslint-windows-386",
    arm: "fslint-windows-arm",
    arm64: "fslint-windows-arm64",
    x64: "fslint-windows-amd64",
  },
  darwin: {
    arm64: "fslint-darwin-arm64",
    amd64: "fslint-darwin-amd64",
  },
  linux: {
    ia32: "fslint-linux-386",
    arm: "fslint-linux-arm",
    arm64: "fslint-linux-arm64",
    x64: "fslint-linux-amd64",
    mips: "fslint-linux-mips",
    mipsel: "fslint-linux-mipsel",
    mips64: "fslint-linux-mips64",
    mips64el: "fslint-linux-mips64el",
  },
  freebsd: {
    ia32: "fslint-freebsd-386",
    arm: "fslint-freebsd-arm",
    arm64: "fslint-freebsd-arm64",
    amd64: "fslint-freebsd-amd64",
  },
  openbsd: {
    ia32: "fslint-openbsd-386",
    arm: "fslint-openbsd-arm",
    arm64: "fslint-openbsd-arm64",
    amd64: "fslint-openbsd-amd64",
  },
};

const archMap = platformMap[platform];

if (!archMap) throw ERR_NOT_SUPPORT;

const prebuildPackageName = archMap[arch];

if (!prebuildPackageName) throw ERR_NOT_SUPPORT;

const binaryPackageDir = path.dirname(
  require.resolve(`@axetroy/${prebuildPackageName}/package.json`)
);

const executableFileName = "fslint" + (platform === "win32" ? ".exe" : "");

const executableFilePath = path.join(binaryPackageDir, executableFileName);

/**
 *
 * @param {Array<string>} argv
 * @param {SpawnOptionsWithoutStdio} [spawnOptions]
 * @returns
 */
function exec(argv, spawnOptions = {}) {
  const ps = spawn(executableFilePath, argv, {
    ...spawnOptions,
    stdout: "piped",
  });

  return ps;
}

/**
 * @param {Object} params0
 * @param {string} params0.config The config file path
 * @param {number} [params0.maxError] The max error
 * @returns {Promise<any>}
 */
function fslint({ config, maxError }) {
  const args = ["--json", "--no-color", "--config", config];

  if (maxError) {
    args.push("--max-error");
    args.push(maxError);
  }

  const ps = exec(args, {
    stdout: "pipe",
    stderr: "pipe",
  });

  let stdout = Buffer.from("");
  let stderr = Buffer.from("");

  ps.stdout.on("data", (/** @type {Buffer} */ buf) => {
    stdout = Buffer.concat(stdout, buf);
  });

  ps.stderr.on("data", (/** @type {Buffer} */ buf) => {
    stderr = Buffer.concat(stderr, buf);
  });

  return new Promise((resolve, reject) => {
    ps.on("exit", (code) => {
      if (code === 0) {
        const output = stdout.toString("utf-8").trim();

        try {
          resolve(JSON.parse(output));
        } catch (err) {
          reject(err);
        }
      } else {
        reject(new Error(`fslint error: \n${stderr.toString("utf-8").trim()}`));
      }
    });
  });
}

module.exports = fslint;
module.exports.fslint = fslint;
module.exports.exec = exec;

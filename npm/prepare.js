const fs = require("fs");
const path = require("path");

const ref = process.env.GIT_REF; // refs/tags/v1.0.0

const arr = ref.split("/");
const version = arr[arr.length - 1].replace(/^v/, "");

console.log(`prepare publish to npm for: ${version}`);

const packages = fs
  .readdirSync(__dirname)
  .filter((v) => v.startsWith("fslint-"))
  .concat(["fslint"]);

for (const pkgName of packages) {
  const pkgPath = path.join(__dirname, pkgName, "package.json");

  const pkg = require(pkgPath);

  pkg.version = version;

  if (pkg.optionalDependencies) {
    for (const subDeps in pkg.optionalDependencies) {
      if (subDeps.startsWith("fslint-")) {
        pkg.optionalDependencies[subDeps] = version;
      }
    }
  }

  fs.writeFileSync(pkgPath, JSON.stringify(pkg, null, 2));

  if (pkgName.startsWith("fslint-")) {
    const fileMap = {
      "fslint-darwin-arm64": "fslint_darwin_arm64",
      "fslint-darwin-amd64": "fslint_darwin_amd64_v1",
      "fslint-linux-arm": "fslint_linux_armv_7",
      "fslint-linux-amd64": "fslint_linux_amd64_v1",
      "fslint-linux-arm64": "fslint_linux_arm64",
      "fslint-windows-386": "fslint_windows_386",
      "fslint-windows-amd64": "fslint_windows_amd64_v1",
      "fslint-windows-arm": "fslint_windows_arm_7",
      "fslint-windows-arm64": "fslint_windows_arm64",
    };

    if (pkgName in fileMap === false)
      throw new Error(`Can not found prebuild file for package '${pkgName}'`);

    const distFolder = fileMap[pkgName];

    const executableFileName =
      "fslint" + (pkgName.indexOf("windows") > -1) ? ".exe" : "";

    const executableFilePath = path.join(
      __dirname,
      "..",
      "dist",
      distFolder,
      executableFileName
    );

    fs.copyFileSync(
      executableFilePath,
      path.join(__dirname, pkgName, executableFileName)
    );
  }
}

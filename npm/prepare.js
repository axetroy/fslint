const fs = require("fs");
const path = require("path");

const ref = process.env.GIT_REF; // refs/tags/v1.0.0

const arr = ref.split("/");
const version = arr[arr.length - 1].replace(/^v/, "");

console.log(`prepare publish to npm for: ${version}`);

const subPackages = fs
  .readdirSync(__dirname)
  .filter((v) => v.startsWith("fslint-"));

for (const subPkg of subPackages) {
  const pkgPath = path.join(__dirname, subPkg, "package.json");

  const pkg = require(pkgPath);

  pkg.version = version;

  fs.writeFileSync(pkgPath, JSON.stringify(pkg, null, 2));
}

const mainPkgPath = path.join(__dirname, "fslint", "package.json");

const mainPkg = require(mainPkgPath);

mainPkg.version = version;

for (const subDeps in mainPkg.optionalDependencies) {
  if (subDeps.startsWith("fslint-")) {
    mainPkg.optionalDependencies[subDeps] = version;
  }
}

fs.writeFileSync(mainPkgPath, JSON.stringify(mainPkg, null, 2));

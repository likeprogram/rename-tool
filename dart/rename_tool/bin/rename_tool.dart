import 'dart:io';

import 'package:args/args.dart';
import 'package:path/path.dart';

const _argDir = "dir";
void main(List<String> arguments) async {
  var argParser = ArgParser();
  argParser.addOption(_argDir, abbr: "d", callback: (dir) => print("Got dir: $dir"));
  var results = argParser.parse(arguments);
  var strDir = current;
  if (results.wasParsed(_argDir) && Directory(results[_argDir]).existsSync()) {
    strDir = results[_argDir];
  }

  var fileNamePat = RegExp(r".+_(bin.+\.dat)");
  var dir = Directory(strDir);
  var dirList = dir.list();
  await for (final FileSystemEntity f in dirList) {
    if (f is File) {
      // print("Found: ${f.path}");
      print("base name: ${basename(f.path)}");
      for (var match in fileNamePat.allMatches(basename(f.path))) {
        // print("group count: ${match.groupCount}");
        if (match.groupCount == 1) {
          var dest = Directory(join(f.parent.path, "dest"));
          if (!dest.existsSync()) {
            dest.createSync();
          }
          // var newFile = await f.rename(join(dest.path, match.group(1)));
          var newFile = await f.copy(join(dest.path, match.group(1)));
          print("after rename, file: ${newFile.path}");
        }
      }
    }
  }
}

CLI tool that prints out metadata stored in [Info.plist](https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/PropertyLists/UnderstandXMLPlist/UnderstandXMLPlist.html) inside an [IPA file](https://www.loc.gov/preservation/digital/formats/fdd/fdd000593.shtml).

Example usage:

```
user@dev1:~/ipaparser$ ./ipaparser --ipa Marvin_Classic_eBook_reader_for_epub_2.9.1.ipa 
Parsed IPA data:
DTCompiler: com.apple.compilers.llvm.clang.1_0
DTSDKBuild: 13E230
UIFileSharingEnabled: true
UIDeviceFamily: [1 2]
CFBundleSignature: ????
DTPlatformVersion: 9.3
DTSDKName: iphoneos9.3
[..]
```


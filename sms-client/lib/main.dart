import 'package:flutter/material.dart';
import 'package:flutter/foundation.dart' show kIsWeb;
import 'app.dart';

void main() {
  // 在 Web 平台强制使用 HTML 渲染器
  if (kIsWeb) {
    // 注意：在新版 Flutter 中，渲染器选择已经自动化
    // 如果仍然有 CanvasKit 问题，需要在 web/index.html 中配置
  }
  
  runApp(const MyApp());
}

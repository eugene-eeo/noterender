!function(){
  var set = document.getElementsByTagName('code');
  var pat = /(?:^\$\$\s*([\s\S]+)\s*\$\$$)|(?:^\$\s*([\s\S]+)\s*\$$)/;

  function extractTeX(text) {
    var match = text.match(pat);
    if (!match) return null;
    return {
      TeX: match[1] || match[2],
      displayMode: match[1] ? true : false,
    };
  }

  for (var i = set.length; i--;) {
    var code = set[i];
    var info = extractTeX(code.textContent);
    if (!info) continue;
    katex.render(info.TeX, code, {
      displayMode: info.displayMode,
      throwOnError: false,
    });
    code.classList.add('has-jax');
  }
}();

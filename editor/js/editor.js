!(function() {
	const editor = CodeMirror.fromTextArea(document.querySelector("#editor"), {
		node: "javascript",
		lineNumbers: true,
	});
	editor.save();


	return true;
})()

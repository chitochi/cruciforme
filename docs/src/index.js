import hljs from 'highlight.js/lib/core'
import xml from 'highlight.js/lib/languages/xml'

hljs.configure({
	tabReplace: '  '
})

hljs.registerLanguage('xml', xml)
hljs.initHighlightingOnLoad()

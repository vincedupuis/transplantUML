# W3C SCXML 1.0 XML Schema

Vendored copy of the normative SCXML schema, used by `TestSCXMLValidatesAgainstW3CSchema`
(`internal/scxml/schema_test.go`) to check both the fixtures we parse and the documents our
emitter writes against the spec rather than against ourselves.

Retrieved 2026-09-21 from <https://www.w3.org/2011/04/SCXML/>:

    scxml.xsd                  scxml-attribs.xsd        scxml-datatypes.xsd
    scxml-module-core.xsd      scxml-contentmodels.xsd  scxml-copyright.xsd
    scxml-module-data.xsd      scxml-module-external.xsd

and `xml.xsd` from <https://www.w3.org/2001/xml.xsd>.

## Local modification

`scxml.xsd` and `scxml-attribs.xsd` both import the XML namespace by absolute URL:

    schemaLocation="http://www.w3.org/2001/xml.xsd"

Both were rewritten to `schemaLocation="xml.xsd"` so validation resolves against the copy in this
directory. Without it xmllint reaches out to w3.org, which makes the test need a network; patching
only one of the two leaves the other fetching remotely and emits a duplicate-import warning.

Nothing else was changed. Re-fetch with the same rewrite if the schema is ever updated.

## License

Distributed under the [W3C Software and Document License](https://www.w3.org/copyright/software-license-2023/).
Copyright (c) World Wide Web Consortium. See `scxml-copyright.xsd` for the notice shipped with the schema.

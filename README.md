<a href="https://codeclimate.com/github/Shelestov7/TextSearchAPI/maintainability"><img src="https://api.codeclimate.com/v1/badges/cf072260977d5e8bb9be/maintainability" /></a>

# TextSearchAPI

Simple RestAPI with single endpoint.

Api provide only one endpoint bind on `/api/v0.1/search/`

Full path `http://host:port/api/v0.1/search/{searchWord}`

#### Endpoint response example:

`{"wordFound":true,"numOccurrences":2,"lineOccurrences":[13]}`

## Project set up
<ol>
  <li><p>Clone the repository from GitHub.com</p>
<p><code>git clone https://github.com/Shelestov7/TextSearchAPI.git</code></p></li>
<li><p>Build docker image</p>
<p><code>docker build -t app-name .</code></p></li>
<li><p>And run</p>
<p><code>docker run -p 8080:8080 my-go-app</code></p></li>
</ol>





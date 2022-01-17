$(function() {
  // Divs for pages where you are viewing a profile and your own.
  const following = $("span.user-following-container");
  const editPage = $("div.mb-3");

  // Try to extract the user. If the path length not 2 exit out.
  const pathname = window.location.pathname;
  let explodedPath = pathname.split("/");
  console.log(explodedPath);
  if (explodedPath.length !== 2) {
    return
  }
  const user = explodedPath[1];

  // Try to find the correct container for the gist
  const followDiv = following[following.length-1];
  const editDiv = editPage[2];

  let lineBreak = document.createElement("br");

  if (user) {
    if (followDiv) {
      $(followDiv).append(lineBreak);
      $(lineBreak).after(`
        <span class="follow d-block">
          <a class="btn btn-block" href="https://gist.github.com/${user}">Gists</a>
        </span>
      `);
    } else if(editDiv) {
      $(editDiv).append(lineBreak);
      $(lineBreak).after(`
        <span class="follow d-block">
          <a class="btn btn-block" href="https://gist.github.com/${user}">Gists</a>
        </span>
      `);
    }
  }
});
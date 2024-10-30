(function NavToggle() {
  addToggleListener("nav-toggle", "nav-menu");
  addToggleListener("docs-nav-mobile-toggle", "docs-nav-mobile-menu");

  function addToggleListener(btnId, navId) {
    const navBtn = document.getElementById(btnId);
    const navMenu = document.getElementById(navId);
    if (!navBtn || !navMenu) {
      return;
    }

    const navIcon = navBtn.getElementsByClassName("toggle-icon")[0];
    const navLinks = navMenu.querySelectorAll("a");
    let expanded = false;

    navBtn.addEventListener("click", () => {
      toggleNav();
    });

    function toggleNav() {
      expanded = !expanded;
      navMenu.classList.toggle("hidden");
      navMenu.classList.toggle("block");
      navMenu.setAttribute("aria-hidden", !expanded);
      navBtn.setAttribute("aria-expanded", expanded);
      if (navIcon) {
        navIcon.classList.toggle("rotate-90");
      }

      navLinks.forEach((link) => {
        link.tabIndex = expanded ? 0 : -1;
      });
      if (expanded) navLinks[0].focus();
    }
  }
})();

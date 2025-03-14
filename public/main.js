(function navToggle() {
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

(function registerFadeInElements() {
  const elements = document.querySelectorAll("[data-fade-in]");
  if (elements.length === 0) return;
  const staggerDelay = 150;
  const duration = 800;
  const intersectionObserver = new IntersectionObserver(
    (entries, observer) => {
      entries.forEach((entry, i) => animateEntry(entry, i, observer));
    },
    {
      root: null,
      threshold: 0.1,
    },
  );
  elements.forEach((element) => {
    element.classList.add("transition");
    element.classList.add("duration-800");
    intersectionObserver.observe(element);
  });
  const animateEntry = (entry, index, observer) => {
    const element = entry.target;
    if (!entry.isIntersecting) {
      element.classList.add("opacity-0");
      element.classList.add("translate-y-10");
      return;
    }
    observer.unobserve(element);
    setTimeout(() => {
      element.classList.remove("opacity-0");
      element.classList.remove("translate-y-10");
      removeTransitionAfterAnimated(element);
    }, index * staggerDelay);
  };
  const removeTransitionAfterAnimated = (element) => {
    setTimeout(() => {
      element.classList.remove("transition");
      element.classList.remove("duration-800");
    }, duration);
  };
})();

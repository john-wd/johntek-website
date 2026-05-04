(function () {
  document.addEventListener("DOMContentLoaded", () => {
    const nav = document.querySelector(".c-nav");
    const toggle = document.querySelector(".c-nav__toggle");

    if (!nav || !toggle) return;

    // mobile hamburger menu listener
    toggle.addEventListener("click", () => {
      const isOpen = nav.classList.toggle("is-open");
      toggle.setAttribute("aria-expanded", String(isOpen));
      toggle.setAttribute(
        "aria-label",
        isOpen ? "Close navigation menu" : "Open navigation menu"
      );
    });

    // scroll-spy to make navbar transparent with
    var cls = 'is-scrolled';
    function check() {
      if (window.scrollY > 50) nav.classList.add(cls);
      else nav.classList.remove(cls);
    }
    window.addEventListener('scroll', check, { passive: true });
    check();
  });
})();
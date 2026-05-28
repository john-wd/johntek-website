function setupNavigation() {
  const nav = document.querySelector(".c-nav");
  const toggle = document.querySelector(".c-nav__toggle");

  if (!nav || !toggle) return;

  // mobile hamburger menu listener
  toggle.addEventListener("click", () => {
    const isOpen = nav.classList.toggle("is-open");
    toggle.setAttribute("aria-expanded", String(isOpen));
    toggle.setAttribute(
      "aria-label",
      isOpen ? "Close navigation menu" : "Open navigation menu",
    );
  });

  // scroll-spy to make navbar transparent with
  var cls = "is-scrolled";
  function check() {
    if (window.scrollY > 50) nav.classList.add(cls);
    else nav.classList.remove(cls);
  }
  window.addEventListener("scroll", check, { passive: true });
  check();
}

function setupCarousel() {
  document.querySelectorAll("[data-carousel]").forEach((carousel) => {
    const track = carousel.querySelector(".c-carousel__track");
    const prev = carousel.querySelector("[data-carousel-prev]");
    const next = carousel.querySelector("[data-carousel-next]");

    let isPaused = false;
    let intervalId = null;

    const getScrollAmount = () => {
      const card = track.querySelector(".c-carousel__item");
      if (!card) return 0;

      const gap = parseFloat(getComputedStyle(track).gap) || 0;
      return card.offsetWidth + gap;
    };

    const scrollByCard = (direction = 1) => {
      const amount = getScrollAmount();
      if (!amount) return;

      const nearEnd =
        track.scrollLeft + track.clientWidth >= track.scrollWidth - amount / 2;

      if (direction > 0 && nearEnd) {
        track.scrollTo({
          left: 0,
          behavior: "smooth",
        });
        return;
      }

      track.scrollBy({
        left: direction * amount,
        behavior: "smooth",
      });
    };

    const startAutoplay = () => {
      stopAutoplay();

      intervalId = setInterval(() => {
        if (!isPaused) {
          scrollByCard(1);
        }
      }, 5000);
    };

    const stopAutoplay = () => {
      if (intervalId) {
        clearInterval(intervalId);
        intervalId = null;
      }
    };

    prev?.addEventListener("click", () => scrollByCard(-1));
    next?.addEventListener("click", () => scrollByCard(1));

    carousel.addEventListener("mouseenter", () => {
      isPaused = true;
    });
    carousel.addEventListener("mouseleave", () => {
      isPaused = false;
    });
    carousel.addEventListener("focusin", () => {
      isPaused = true;
    });
    carousel.addEventListener("focusout", () => {
      isPaused = false;
    });

    startAutoplay();
  });
}

function setupMobileShare() {
  document.addEventListener("click", async (event) => {
      const button = event.target.closest(".c-share-native");
      if (!button) return;
  
      const title = button.dataset.title;
      const url = button.dataset.url;
  
      if (navigator.share) {
        try {
          await navigator.share({ title, url });
        } catch (err) {
          // User cancelled share sheet; ignore.
        }
      } else {
        await navigator.clipboard.writeText(url);
        button.textContent = "Copied";
        setTimeout(() => {
          button.textContent = "Share";
        }, 1500);
      }
    });

}

(function () {
  document.addEventListener("DOMContentLoaded", () => {
    setupNavigation();
    setupCarousel();
    setupMobileShare();
  });
})();

var MyScroll = "";
(function (window, document, $, undefined) {
  "use strict";
  var isMobile =
    /Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Nokia|Opera Mini/i.test(
      navigator.userAgent
    )
      ? !0
      : !1;
  var Scrollbar = window.Scrollbar;
  var Init = {
    i: function (e) {
      Init.s();
      Init.methods();
    },
    s: function (e) {
      (this._window = $(window)),
        (this._document = $(document)),
        (this._body = $("body")),
        (this._html = $("html"));
    },
    methods: function (e) {
      Init.w();
      Init.BackToTop();
      Init.preloader();
      Init.header();
      Init.slick();
      Init.categoryToggle();
      Init.filterSearch();
      Init.passwordIcon();
      Init.formValidation();
      Init.contactForm();
      Init.quantityHandle();
      Init.filterToggle();
      Init.priceRangeSlider();
      Init.billingAddress();
      Init.checkBoxes();
      Init.toggles();
      Init.brandsblock();
      Init.dropdown();
      Init.quantityHandle();
      Init.showReview();
    },

    BackToTop: function () {
      var scrollToTopBtn = document.querySelector(".scrollToTopBtn");
      var rootElement = document.documentElement;
      function handleScroll() {
        var scrollTotal = rootElement.scrollHeight - rootElement.clientHeight;
        if (rootElement.scrollTop / scrollTotal > 0.05) {
          scrollToTopBtn.classList.add("showBtn");
        } else {
          scrollToTopBtn.classList.remove("showBtn");
        }
      }
      function scrollToTop() {
        rootElement.scrollTo({ top: 0, behavior: "smooth" });
      }
      scrollToTopBtn.addEventListener("click", scrollToTop);
      document.addEventListener("scroll", handleScroll);
    },

    preloader: function () {
      setTimeout(function () {
        $("#preloader").hide("slow");
      }, 3600);
    },

    showReview: function () {
      $(".review-btn").on("click", function () {
        $(".review-btn").removeClass("te-button");
        var id = $(this).attr("data-atr");
        $(this).addClass("te-button");

        // Hide all review blocks slowly
        $(".review-block").hide("slow");

        // Show the selected review block slowly
        $("#" + id).show("slow");
      });
    },
    teamMemberShow: function (e) {
      $(".member").on("click", function () {
        var id = $(this).attr("id");
        $(".member").removeClass("active");
        $(this).addClass("active");
        $(".member-details").hide("slow");
        $("." + id).show("slow");
      });
    },

    serviceShow: function (e) {
      $(".service_title").on("click", function () {
        var id = $(this).attr("id");
        $(".service_title").removeClass("active");
        $(this).addClass("active");
        $(".service-detail").hide("slow");
        $("." + id).show("slow");
      });
    },

    w: function (e) {
      if (isMobile) {
        $("body").addClass("is-mobile");
      }
    },

    header: function () {
      function dynamicCurrentMenuClass(selector) {
        let FileName = window.location.href.split("/").reverse()[0];
        selector.find("li").each(function () {
          let anchor = $(this).find("a");
          if ($(anchor).attr("href") == FileName) {
            $(this).addClass("current");
          }
        });
        selector.children("li").each(function () {
          if ($(this).find(".current").length) {
            $(this).addClass("current");
          }
        });
        if ("" == FileName) {
          selector.find("li").eq(0).addClass("current");
        }
      }
      if ($(".main-menu__list").length) {
        let mainNavUL = $(".main-menu__list");
        dynamicCurrentMenuClass(mainNavUL);
      }
      if ($(".main-menu__nav").length && $(".mobile-nav__container").length) {
        let navContent = document.querySelector(".main-menu__nav").innerHTML;
        let mobileNavContainer = document.querySelector(
          ".mobile-nav__container"
        );
        mobileNavContainer.innerHTML = navContent;
      }
      if ($(".sticky-header__content").length) {
        let navContent = document.querySelector(".main-menu").innerHTML;
        let mobileNavContainer = document.querySelector(
          ".sticky-header__content"
        );
        mobileNavContainer.innerHTML = navContent;
      }
      if ($(".mobile-nav__container .main-menu__list").length) {
        let dropdownAnchor = $(
          ".mobile-nav__container .main-menu__list .dropdown > a"
        );
        dropdownAnchor.each(function () {
          let self = $(this);
          let toggleBtn = document.createElement("BUTTON");
          toggleBtn.setAttribute("aria-label", "dropdown toggler");
          toggleBtn.innerHTML = "<i class='fa fa-angle-down'></i>";
          self.append(function () {
            return toggleBtn;
          });
          self.find("button").on("click", function (e) {
            e.preventDefault();
            let self = $(this);
            self.toggleClass("expanded");
            self.parent().toggleClass("expanded");
            self.parent().parent().children("ul").slideToggle();
          });
        });
      }
      if ($(".mobile-nav__toggler").length) {
        $(".mobile-nav__toggler").on("click", function (e) {
          e.preventDefault();
          $(".mobile-nav__wrapper").toggleClass("expanded");
          $("body").toggleClass("locked");
        });
      }
      $(window).on("scroll", function () {
        if ($(".stricked-menu").length) {
          var headerScrollPos = 130;
          var stricky = $(".stricked-menu");
          if ($(window).scrollTop() > headerScrollPos) {
            stricky.addClass("stricky-fixed");
          } else if ($(this).scrollTop() <= headerScrollPos) {
            stricky.removeClass("stricky-fixed");
          }
        }
      });
    },

    smoothScrollbar: function () {
      if ($("body").hasClass("tt-smooth-scroll")) {
        if (!isMobile) {
          class AnchorPlugin extends Scrollbar.ScrollbarPlugin {
            static pluginName = "anchor";
            onHashChange = () => {
              $(".header-menu").animate({ height: "toggle" });
              this.jumpToHash(window.location.hash);
            };
            jumpToHash = (hash) => {
              if (!hash) {
                return;
              }
              const { scrollbar } = this;
              scrollbar.containerEl.scrollTop = 0;
              const target = document.querySelector(hash);
              if (target) {
                scrollbar.scrollIntoView(target, {
                  offsetTop:
                    parseFloat(target.getAttribute("data-offset")) || 0,
                });
              }
            };
            onInit() {
              this.jumpToHash(window.location.hash);
              window.addEventListener("hashchange", this.onHashChange);
            }
            onDestory() {
              window.removeEventListener("hashchange", this.onHashChange);
            }
          }
          Scrollbar.use(AnchorPlugin);
          const scrollbar = Scrollbar.init(
            document.querySelector("#scroll-container"),
            { damping: 0.1, renderByPixel: !0, continuousScrolling: !0 }
          );
          $("input[type=number]").on("focus", function () {
            $(this).on("wheel", function (e) {
              e.stopPropagation();
            });
          });
          const backToTopButton = document.getElementById("back-to-top");
          scrollbar.addListener(({ offset }) => {
            if (offset.y > 300) {
              backToTopButton.style.display = "block";
            } else {
              backToTopButton.style.display = "none";
            }
          });
          backToTopButton.addEventListener("click", () => {
            scrollbar.scrollTo(0, 0, 500);
          });
        }
      }
    },

    slick: function () {
      if ($(".gallery-slider").length) {
        $(".gallery-slider").slick({
          slidesToShow: 1,
          infinite: !0,
          autoplay: true,
          dots: false,
          draggable: true,
          arrows: false,
          lazyLoad: "linear",
          speed: 800,
          autoplaySpeed: 1400,
          responsive: [
            { breakpoint: 1025, settings: { slidesToShow: 1 } },
            { breakpoint: 992, settings: { slidesToShow: 1 } },
            { breakpoint: 576, settings: { slidesToShow: 1 } },
          ],
        });
      }
      if ($(".our-journey-slider").length) {
        $(".our-journey-slider").slick({
          slidesToShow: 4,
          variableWidth: false,
          infinite: !0,
          autoplay: false,
          dots: false,
          centerMode: false,
          draggable: true,
          arrows: false,
          lazyLoad: "linear",
          speed: 800,
          autoplaySpeed: 3000,
          responsive: [
            { breakpoint: 1499, settings: { slidesToShow: 3 } },
            { breakpoint: 1025, settings: { slidesToShow: 1 } },
            { breakpoint: 992, settings: { slidesToShow: 1 } },
            { breakpoint: 576, settings: { slidesToShow: 1 } },
          ],
        });
      }
      if ($(".upcoming-match-slider").length) {
        $(".upcoming-match-slider").slick({
          slidesToShow: 4,
          variableWidth: false,
          infinite: !0,
          autoplay: false,
          dots: false,
          centerMode: false,
          draggable: true,
          arrows: false,
          lazyLoad: "linear",
          speed: 800,
          autoplaySpeed: 2000,
          responsive: [
            { breakpoint: 1499, settings: { slidesToShow: 3 } },
            { breakpoint: 1025, settings: { slidesToShow: 2 } },
            { breakpoint: 992, settings: { slidesToShow: 1 } },
            { breakpoint: 576, settings: { slidesToShow: 1 } },
          ],
        });
      }
      if ($(".our-community-slider").length) {
        $(".our-community-slider").slick({
          slidesToShow: 2,
          variableWidth: false,
          infinite: !0,
          autoplay: false,
          dots: false,
          centerMode: false,
          draggable: true,
          arrows: false,
          lazyLoad: "linear",
          speed: 800,
          autoplaySpeed: 2000,
          responsive: [
            { breakpoint: 1399, settings: { slidesToShow: 1 } },
            { breakpoint: 992, settings: { slidesToShow: 1 } },
            { breakpoint: 576, settings: { slidesToShow: 1 } },
          ],
        });
      }
      if ($(".brand-slider").length) {
        $(".brand-slider").slick({
          infinite: true,
          slidesToShow: 6,
          arrows: false,
          autoplay: true,
          cssEase: "linear",
          autoplaySpeed: 0,
          speed: 8000,
          pauseOnFocus: false,
          pauseOnHover: false,
          responsive: [
            {
              breakpoint: 1699,
              settings: {
                slidesToShow: 5,
              },
            },
            {
              breakpoint: 1599,
              settings: {
                slidesToShow: 4,
              },
            },
            {
              breakpoint: 769,
              settings: {
                slidesToShow: 3,
              },
            },
            {
              breakpoint: 576,
              settings: {
                slidesToShow: 2,
              },
            },
            {
              breakpoint: 450,
              settings: {
                slidesToShow: 1,
              },
            },
          ],
        });
      }

      if ($(".product-slider").length) {
        $(".product-slider").slick({
          slidesToShow: 1,
          slidesToScroll: 1,
          arrows: false,
          fade: true,
          asNavFor: ".product-slider-asnav",
        });
      }
      if ($(".product-slider-asnav").length) {
        $(".product-slider-asnav").slick({
          slidesToShow: 4,
          slidesToScroll: 1,
          asNavFor: ".product-slider",
          dots: false,
          arrows: false,
          centerMode: false,
          variableWidth: true,
          focusOnSelect: true,
        });
      }

      $(".btn-prev").click(function () {
        var $this = $(this).attr("data-slide");
        $("." + $this).slick("slickPrev");
      });
      $(".btn-next").click(function () {
        var $this = $(this).attr("data-slide");
        $("." + $this).slick("slickNext");
      });
    },

    // Quantity Controller
    quantityHandle: function () {
      $(".decrement").on("click", function () {
        var qtyInput = $(this).closest(".quantity-wrap").children(".number");
        var qtyVal = parseInt(qtyInput.val());
        if (qtyVal > 0) {
          qtyInput.val(qtyVal - 1);
        }
      });
      $(".increment").on("click", function () {
        var qtyInput = $(this).closest(".quantity-wrap").children(".number");
        var qtyVal = parseInt(qtyInput.val());
        qtyInput.val(parseInt(qtyVal + 1));
      });
    },

    // Filter Toggle Button
    filterToggle: function () {
      if ($(".category-block").length) {
        $(".category-block .title").on("click", function (e) {
          var count = $(this).data("count");
          if (
            $(".category-block.box-" + count + " .content-block").is(":visible")
          ) {
            $(".category-block.box-" + count + " span i").removeClass(
              "fa-horizontal-rule"
            );
            $(".category-block.box-" + count + " span i").addClass("fa-plus");
            $(".category-block.box-" + count + " .content-block").hide("slow");
          } else {
            $(".category-block.box-" + count + " span i").removeClass(
              "fa-plus"
            );
            $(".category-block.box-" + count + " span i").addClass(
              "fa-horizontal-rule"
            );
            $(".category-block.box-" + count + " .content-block").show("slow");
          }
        });
      }
      if ($(".toggle-sidebar").length) {
        $(".shop-filter").on("click", function () {
          $(".toggle-sidebar").animate({ left: "0" }, 300);
          $(".overlay").fadeIn(300);
        });

        $(".overlay").on("click", function () {
          $(".toggle-sidebar").animate({ left: "-400px" }, 300);
          $(this).fadeOut(300);
        });
      }
    },

    // Form Validation
    priceRangeSlider: function () {
      const priceGap = 1000;

      $(".price-input input").on("input", function () {
        let minPrice = parseInt($(".price-input .input-min").val()),
          maxPrice = parseInt($(".price-input .input-max").val());

        if (
          maxPrice - minPrice >= priceGap &&
          maxPrice <= $(".range-input .range-max").attr("max")
        ) {
          if ($(this).hasClass("input-min")) {
            $(".range-input .range-min").val(minPrice);
            $(".slider .progress").css(
              "left",
              (minPrice / $(".range-input .range-min").attr("max")) * 100 + "%"
            );
          } else {
            $(".range-input .range-max").val(maxPrice);
            $(".slider .progress").css(
              "right",
              100 -
                (maxPrice / $(".range-input .range-max").attr("max")) * 100 +
                "%"
            );
          }
        }
      });

      $(".range-input input").on("input", function () {
        let minVal = parseInt($(".range-input .range-min").val()),
          maxVal = parseInt($(".range-input .range-max").val());

        if (maxVal - minVal < priceGap) {
          if ($(this).hasClass("range-min")) {
            $(".range-input .range-min").val(maxVal - priceGap);
          } else {
            $(".range-input .range-max").val(minVal + priceGap);
          }
        } else {
          $(".price-input .input-min").val(minVal);
          $(".price-input .input-max").val(maxVal);
          $(".slider .progress").css(
            "left",
            (minVal / $(".range-input .range-min").attr("max")) * 100 + "%"
          );
          $(".slider .progress").css(
            "right",
            100 -
              (maxVal / $(".range-input .range-max").attr("max")) * 100 +
              "%"
          );
        }
      });
    },

    // Toggle CheckBoxes
    checkBoxes: function () {
      $(".sub-checkboxes").hide();
      $(".arrow-block").click(function () {
        var subCheckboxes = $(this).next(".sub-checkboxes");
        var chevronIcon = $(this).find("i");
        subCheckboxes.slideToggle("fast");
        chevronIcon.toggleClass("fa-chevron-down fa-chevron-up");
      });
      $(".check-block, .sub-check-box").click(function (event) {
        event.stopPropagation();
      });

      if ($(".customer-container").length) {
        $(".signin-button").click(function () {
          $(".sign-form").slideToggle();
        });
      }
    },

    
    // CheckOut Same Billing Address
    billingAddress: function () {
      if ($("#shipAddress").length) {
        $('.billing-address').hide();
        $('#shipAddress').change(function () {
          if ($(this).is(':unchecked')) {
            $('.billing-address').hide("slow");
          } else {
            $('.billing-address').show("slow");
          }
        });
      }

      if ($("#open-menu-all").length) {
        $('.all-category-list').hide();

        $('#open-menu-all').on('click', function (event) {
          event.stopPropagation(); // Prevent click from propagating to document

          if ($('.all-category-list').is(':visible')) {
            $('.all-category-list').slideUp("slow");
            $('body').removeClass('sidebar-active');
            $('.all-navigator i').removeClass('fa-times').addClass('fa-bars');
          } else {
            $('.all-category-list').slideDown("slow");
            setTimeout(function () {
              $('body').addClass('sidebar-active');
              $('.all-navigator i').removeClass('fa-bars').addClass('fa-times');
            }, 500);
          }
        });

        $(document).on('click', function () {
          if ($('body').hasClass('sidebar-active')) {
            $('.all-category-list').slideUp("slow");
            $('body').removeClass('sidebar-active');
            $('.all-navigator i').removeClass('fa-times').addClass('fa-bars');
          }
        });

        // Prevent click on the category list from closing it
        $('.all-category-list').on('click', function (event) {
          event.stopPropagation();
        });
      }
    },


    toggles: function () {
      if ($(".sidebar-widget").length) {
        $(".widget-title-row").on("click", function (e) {
          $(this).find("i").toggleClass("fa-horizontal-rule fa-plus", "slow");
          // $(this).find('i').toggleClass('fa-light fa-regular', 'slow');
          $(this)
            .parents(".sidebar-widget")
            .find(".widget-content-block")
            .animate({ height: "toggle" }, "slow");
        });
      }
      // Wishlist Toggle
      if ($(".wishlist-icon").length) {
        $(".wishlist-icon").on("click", function () {
          $(this).find(".fa-light").toggleClass("fa-solid");
          return false;
        });
      }
    },

    brandsblock: function () {
      document.querySelectorAll(".custom-link").forEach((link) => {
        link.addEventListener("click", function (event) {
          event.preventDefault(); // Prevent the default link behavior

          // Remove 'active' class from all links
          document
            .querySelectorAll(".custom-link")
            .forEach((l) => l.classList.remove("active"));

          // Add 'active' class to the clicked link
          this.classList.add("active");
        });
      });
    },

    // Cart Sidebar
    cartSidebar: function () {
      $(".cart-button").on("click", function () {
        $("#sidebar-cart").css("right", "0");
        $("#sidebar-cart-curtain")
          .fadeIn(0)
          .css("display", "block")
          .animate({ opacity: 1 }, 200); // Smooth fade in the curtain
      });

      $(".close-popup").on("click", function () {
        $("#sidebar-cart").css("right", "-101%");
        $("#sidebar-cart-curtain").animate({ opacity: 0 }, 200, function () {
          $(this).css("display", "none");
        });
      });
    },

    dropdown: function () {
      const selectedAll = document.querySelectorAll(".wrapper-dropdown");

      selectedAll.forEach((selected) => {
        const optionsContainer = selected.children[2];
        const optionsList = selected.querySelectorAll(
          "div.wrapper-dropdown li"
        );

        selected.addEventListener("click", () => {
          let arrow = selected.children[1];

          if (selected.classList.contains("active")) {
            handleDropdown(selected, arrow, false);
          } else {
            let currentActive = document.querySelector(
              ".wrapper-dropdown.active"
            );

            if (currentActive) {
              let anotherArrow = currentActive.children[1];
              handleDropdown(currentActive, anotherArrow, false);
            }

            handleDropdown(selected, arrow, true);
          }
        });

        // update the display of the dropdown
        for (let o of optionsList) {
          o.addEventListener("click", () => {
            selected.querySelector(".selected-display").innerHTML = o.innerHTML;
          });
        }
      });

      // check if anything else ofther than the dropdown is clicked
      window.addEventListener("click", function (e) {
        if (e.target.closest(".wrapper-dropdown") === null) {
          closeAllDropdowns();
        }
      });

      // close all the dropdowns
      function closeAllDropdowns() {
        const selectedAll = document.querySelectorAll(".wrapper-dropdown");
        selectedAll.forEach((selected) => {
          const optionsContainer = selected.children[2];
          let arrow = selected.children[1];

          handleDropdown(selected, arrow, false);
        });
      }

      // open all the dropdowns
      function handleDropdown(dropdown, arrow, open) {
        if (open) {
          arrow.classList.add("rotated");
          dropdown.classList.add("active");
        } else {
          arrow.classList.remove("rotated");
          dropdown.classList.remove("active");
        }
      }
    },

    categoryToggle: function () {
      if ($(".customer-container").length) {
        $(".signin-button").click(function () {
          $(".sign-form").slideToggle();
        });
      }
      if ($(".sidebar").length) {
        $(".shop-filter").on("click", function () {
          $(".toggle-sidebar").toggleClass("active");
          $(".overlay").toggleClass("active");
        });

        // Hide sidebar and overlay when overlay is clicked
        $(".overlay").on("click", function () {
          $(".toggle-sidebar").removeClass("active");
          $(this).removeClass("active");
        });
      }
      if ($("#difrent-ship").length) {
        $("#difrent-ship").on("click", function () {
          $(".box-hide").animate({ height: "toggle" }, 300);
        });
      }
    },

    filterSearch: function () {
      if ($("#searchInput").length) {
        $("#searchInput").on("keyup", function () {
          var value = $(this).val().toLowerCase();
          $(".blogs-block").filter(function () {
            var hasMatch =
              $(this).find(".blog-title").text().toLowerCase().indexOf(value) >
              -1;
            $(this).toggle(hasMatch);
          });
        });
      }
    },
    passwordIcon: function () {
      $("#eye , #eye-icon").click(function () {
        if ($(this).hasClass("fa-eye-slash")) {
          $(this).removeClass("fa-eye-slash");
          $(this).addClass("fa-eye");
          $(".password-input").attr("type", "text");
        } else {
          $(this).removeClass("fa-eye");
          $(this).addClass("fa-eye-slash");
          $(".password-input").attr("type", "password");
        }
      });
    },
    formValidation: function () {
      if ($(".contact-form").length) {
        $(".contact-form").validate();
      }
      if ($(".login-form").length) {
        $(".login-form").validate();
      }
    },
    contactForm: function () {
      $(".contact-form").on("submit", function (e) {
        e.preventDefault();
        if ($(".contact-form").valid()) {
          var _self = $(this);
          _self
            .closest("div")
            .find('button[type="submit"]')
            .attr("disabled", "disabled");
          var data = $(this).serialize();
          $.ajax({
            url: "./assets/mail/contact.php",
            type: "post",
            dataType: "json",
            data: data,
            success: function (data) {
              $(".contact-form").trigger("reset");
              _self.find('button[type="submit"]').removeAttr("disabled");
              if (data.success) {
                document.getElementById("message").innerHTML =
                  "<h5 class='black mt-16 mb-16'>Email Sent Successfully</h5>";
              } else {
                document.getElementById("message").innerHTML =
                  "<h5 class='black mt-16 mb-16'>There is an error</h5>";
              }
              $("#messages").show("slow");
              $("#messages").slideDown("slow");
              setTimeout(function () {
                $("#messages").slideUp("hide");
                $("#messages").hide("slow");
              }, 4000);
            },
          });
        } else {
          return !1;
        }
      });
    },
  };
  Init.i();
})(window, document, jQuery);

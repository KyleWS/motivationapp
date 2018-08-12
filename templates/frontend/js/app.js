"use strict";

window.onload =function() {
   $(".1_scroll").click(function() {
      $("html, body").animate({ scrollTop: $('#features').offset().top }, 1000);
   })
}

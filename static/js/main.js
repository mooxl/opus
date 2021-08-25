var burger = document.querySelector(".navigation__middle--burgerMenu");
var lettuce = document.querySelectorAll(".navigation__middle--burgerMenu-burger");
var navbar = document.querySelector("#navigation");
var toggle = false;
burger.addEventListener('click', toggleF);

function toggleF() {   
    if (Math.max(document.documentElement.clientWidth, window.innerWidth) < 770)
        if (toggle) {
            navbar.style.height = "5em";
            toggle = false;
            lettuce.forEach(element => {
                element.classList.remove("change");
            });
        } else {
            lettuce.forEach(element => {
                element.classList.add("change");
            });
            navbar.style.height = "30rem";
            toggle = true;
        }
}

/* $('a[href*="#"]').not('[href="#"]').not('[href="#0"]').click(function (event) {
    if (location.pathname.replace(/^\//, '') == this.pathname.replace(/^\//, '') && location.hostname == this.hostname) {
        var target = $(this.hash);
        target = target.length ? target : $('[name=' + this.hash.slice(1) + ']');
        if (target.length) {
            event.preventDefault();
            $('html, body').animate({
                scrollTop: target.offset().top
            }, 1, function () {
                var $target = $(target);
                $target.focus();
                if ($target.is(":focus")) {
                    return false;
                } else {
                    $target.attr('tabindex', '-1');
                    $target.focus();
                }
            });
        }
    }
}); */
window.onload = ()=>{
    let menutrigger = document.querySelector(".titlewrapper svg");
    menutrigger.onmouseover = function () {
        let menu = document.querySelector(".menu");
        menutrigger.style.background = "var(--color-primary-light)";
        menutrigger.style.borderRadius = "0.5rem";
        menu.style.display = "flex";
        menu.onmouseleave = ()=>{
            menu.style.display = "none";
            menutrigger.style.background = "none";
        }
    }
};

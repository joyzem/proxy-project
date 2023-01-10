const navbar = document.createElement('nav')
navbar.classList.add('z-depth-0')

const navWrapper = document.createElement('div')
navWrapper.classList.add('nav-wrapper', 'blue')
navbar.appendChild(navWrapper)

const navList = document.createElement('ul')
navList.classList.add('left')
navWrapper.appendChild(navList)

const navLinks = [
    {
        label: 'О проекте',
        href: '/',
    },
    {
        label: 'Доверенности',
        href: '/proxy/',
    },
    {
        label: 'Товары',
        href: '/product/',
    },
    {
        label: 'Сотрудники',
        href: '/employee/',
    },
    {
        label: 'Контрагенты',
        href: '/customer/',
    },
    {
        label: 'Организации',
        href: '/organization/',
    },
    {
        label: 'Счета',
        href: '/account/',
    },
];

navLinks.forEach(link => {
    const navItem = document.createElement('li')
    navItem.classList.add('nav-item')
    navList.appendChild(navItem)

    const navLink = document.createElement('a')
    navLink.textContent = link.label
    navLink.href = link.href
    navItem.appendChild(navLink)
});

const navbarContainer = document.createElement('div')
navbarContainer.classList.add('navbar-fixed')
navbarContainer.appendChild(navbar)

// Append the navbar to the DOM
document.querySelector('header').appendChild(navbarContainer)

const currentPath = window.location.pathname

const navItems = document.querySelectorAll('.nav-item');
const isAboutProjectTab = currentPath == '/' || currentPath.includes('/about')
if (isAboutProjectTab) {
    navItems[0].classList.add('active')
} else {
    for (let i = 1; i < navItems.length; i++) {
        item = navItems[i]
        let tabLink = item.querySelector('a').getAttribute('href')
        if (!isAboutProjectTab && currentPath.includes(tabLink)) {
            item.classList.add('active')
        } else {
            item.classList.remove('active')
        }
    }
}

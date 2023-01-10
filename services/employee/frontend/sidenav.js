const sidebar = document.createElement('ul')
sidebar.classList.add('sidenav')
sidebar.id = "sidenav-left"
sidebar.style = "transform: translateX(0px);"

const logoContainer = document.createElement('li')
logoContainer.classList.add('logo-container')
logoContainer.textContent = "Wiki"

const logoIcon = document.createElement('i')
logoIcon.classList.add('material-icons', 'left')
logoIcon.textContent = "language"

logoContainer.appendChild(logoIcon)
sidebar.appendChild(logoContainer)
document.querySelector('header').appendChild(sidebar)

const sidebarLinks = [
    {
        label: 'О сервисе',
        href: '/employee/',
    },
    {
        label: 'Сотрудники',
        href: '/employee/employees',
    },
];

sidebarLinks.forEach(link => {
    const navItem = document.createElement('li')
    navItem.classList.add('bold')
    if (window.location.pathname == link.href) {
        navItem.classList.add('active')
    }

    const itemLink = document.createElement('a')
    itemLink.classList.add('waves-effect')
    itemLink.href = link.href
    itemLink.textContent = link.label

    navItem.appendChild(itemLink)
    sidebar.appendChild(navItem)
})
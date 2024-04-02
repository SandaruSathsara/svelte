import Navbar from "$lib/Components/Navbar.svelte";
import nikelogo from "../assets/nike.png";

export const navbarData = {

    logo: true,
    logosrc: nikelogo,
    logolink: true,
    linkurl: '#/',
    optionallinkText:'velte',
    altText: 'Logo',


    links:[
        {url:'#/',
        displayInNav: true,
        displayInFooter: true,
        linkText: 'Home'
    },
    {url:'#/about',
    displayInNav: true,
    displayInFooter: true,
    linkText: 'About'
    },
    {url:'#/contact',
    displayInNav: true,
    displayInFooter: true,
    linkText: 'Contact'
    },
    {url:'#/help',
    displayInNav: true,
    displayInFooter: true,
    linkText: 'Help'
    },
    {url:'#/notfound',
    displayInNav: false,
    displayInFooter: false,
    linkText: ''
    },
    ]
    
}
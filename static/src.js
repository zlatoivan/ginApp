const domain = 'http://localhost:8080'

const getEl = (id) => {
    return document.getElementById(id)
}

/*const chPassBtnClick = () => {
    fetch(domain + '/user/chPass', {
        method: 'put',
        body: JSON.stringify({
            prevPass: getEl('prevPass').value, // !!!!! .value
            newPass: getEl('newPass').value
        })
    })

    //location.replace(domain + 'chPass-success')
}*/

const deleteAccBtnClick = () => {
    fetch(domain + '/user', {
        method: 'delete',
    })

    location.replace(domain)
}

const showAll = () => {
    fetch(domain + '/user', {
        method: 'delete',
    })
}
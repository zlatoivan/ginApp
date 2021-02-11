const domain = 'http://localhost:8080'

const disBlock = (id) => {
    document.getElementById(id).style.display = "block"
}

const disNone = (id) => {
    document.getElementById(id).style.display = "none"
}

const getEl = (id) => {
    return document.getElementById(id)
}

const addAuthHead = () => {
    return {
        method: 'GET',
        headers: {
            'Authorization': '' + localStorage.getItem('token') // Надо Basic
        }
    }
}

const clickRegistrNav = () => {
    return fetch(domain + "/registr", addAuthHead())
}

// Регистрация
const clickRegistr = () => {
    const username = getEl('username').value
    const password = getEl('password').value
    const controller = new AbortController()
    return fetch(domain + '/user', {
        method: 'POST',
        headers: {
            "Content-Type": "application/json; charset=UTF-8"
        },
        body: JSON.stringify({
            username: username,
            password: password
        }),
        signal: controller.signal
    })
        .then((response) => {
            if (response.status === 409) {
                disBlock('error')
                getEl('error').innerText = "Этот логин занят"
                controller.abort()
            }
            if (response.ok) {
                getEl('author').innerText = username
                return response.text().then((token) => {
                    localStorage.setItem('token', token)
                    disNone('error')
                    //window.location.replace(domain + "/registr-success")
                    return fetch(domain + "/registr-success", addAuthHead())
                })
            }
        })
}

// Авторизация
const clickAuth = () => {
    const username = getEl('username').value
    const password = getEl('password').value
    const controller = new AbortController()
    return fetch(domain + '/user/auth', {
        method: 'POST',
        headers: {
            "Content-Type": "application/json; charset=UTF-8"
        },
        body: JSON.stringify({
            username: username,
            password: password
        }),
        signal: controller.signal
    })
        .then((response) => {
            if (response.status === 403) {
                disBlock('error')
                getEl('error').innerText = "Неверный логин или пароль"
                controller.abort()
            }
            if (response.ok) {
                getEl('author').innerText = username
                return response.text().then((token) => {
                    localStorage.setItem('token', token)
                    disNone('error')
                    return fetch(domain + "/auth-success", addAuthHead())
                })
            }
        })
}

// Выход из учетной записи
const clickLogout = () => {
    localStorage.removeItem('token')
    getEl('author').innerText = "Guest"
    window.location.replace(domain + "/")
}

// Страница изменения пароля
const clickChangePassPage = () => {

}

// Изменение пароля
const clickChangePass = () => {
    const username = getEl('username').value
    const password = getEl('password').value
    return fetch(domain + '/user', {
        method: 'PUT',
        headers: {
            "Content-Type": "application/json; charset=UTF-8"
        },
        body: JSON.stringify({
            username: username,
            password: password
        })
    }).then(() => {})
}

// Вывести всех пользователей
const clickUsers = () => {
    return fetch(domain + '/user', {
        method: 'GET',
        headers: {
            'Authorization': '' + localStorage.getItem('token') // Надо Basic
        }
    })
}





/*.catch((err) => {
    console.log("Error ", err)
})*/
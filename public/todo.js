const viewTodo = function(todos){
    let body = document.getElementById("Todo");
    let list = document.createElement("ol");
    todos.forEach((todo)=>{
        const todoPara = createHtmlTodo(todo);
        list.appendChild(todoPara)
    });
        body.appendChild(list)
};

const loadTodo = function(){
    const xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function() {
        if (xhr.readyState === XMLHttpRequest.DONE) {
            response = JSON.parse(xhr.responseText);
            //console.log(response);
            viewTodo(response)
        }
    };
    xhr.open('GET', '/todo', true);
    xhr.send(null);
};

const createHtmlTodo = function(todo) {
    let todoPara = document.createElement("li");
    let todoElement = document.createElement("h3");
    todoPara.id = todo.id;
    todoElement.innerText = todo.title;
    let todoDescElement = document.createElement("h5");
    todoDescElement.innerText = todo.description;
    todoPara.appendChild(todoElement);
    todoPara.appendChild(todoDescElement);
    return todoPara;
};

const newTodo = function(){
    const title = document.querySelector("#title").value;
    const description = document.querySelector("#description").value;
    const data = {title:title,description:description};
    xhttpRequest("post",update,"create",data)
};

const xhttpRequest = function(type,method,path,data=null){
    const xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function() {
        if (xhr.readyState === XMLHttpRequest.DONE) {
            method()
        }
    };
    xhr.open(type, path, true);
    xhr.setRequestHeader("Content-Type","application/json")
    xhr.send(JSON.stringify(data));
};

const update = function(){
    //console.log("updated")
};
window.onload = loadTodo;

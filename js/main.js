document.getElementById('addTaskButton').addEventListener('click', function () {
    const taskInput = document.getElementById('taskInput');
    const taskText = taskInput.value.trim();
    if (taskText) {
        const taskList = document.getElementById('taskList');
        const li = document.createElement('li');
        li.innerText = taskText;
        
        
        const deleteButton = document.createElement('button');
        deleteButton.innerText = 'Удалить';
        deleteButton.addEventListener('click', function () {
            taskList.removeChild(li);
        });
        
        li.appendChild(deleteButton);
        taskList.appendChild(li);
        taskInput.value = ''; 
    }
});


document.getElementById('feedbackForm').addEventListener('submit', function (event) {
    event.preventDefault();
    
    const name = document.getElementById('name').value;
    const email = document.getElementById('email').value;
    
    const message = document.getElementById('message').value;
    
    alert(`Спасибо за ваше сообщение, ${name}!\nМы свяжемся с вами по адресу ${email}.`);
    
    document.getElementById('feedbackForm').reset();
});


document.getElementById('themeToggleButton').addEventListener('click', function () {
    document.body.classList.toggle('dark-theme');
});


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



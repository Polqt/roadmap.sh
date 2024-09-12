const inputTask = document.getElementById("task-input");
const submitButton = document.getElementById("submit-icon");
const listContainer = document.getElementById("list");
const completeTask = document.getElementById("completed-task");
const incompleteTask = document.getElementById("incomplete-task");

function updateTask() {
    const completedTasks = document.querySelectorAll(".completed").length;
    const incompletedTasks = document.querySelectorAll("li:not(.completed)").length;

    completeTask.textContent = completedTasks;
    incompleteTask.textContent = incompletedTasks;
}

function addTask() {
    const task = inputTask.value.trim();
    const listOfTask = document.createElement("li");
    const hr = document.createElement("hr");

    if (!task) {
        alert("Write down a task.");
        return;
    }

    listOfTask.innerHTML = `
    <label>
        <input type="checkbox">
        <span>${task}</span>
    </label>
    <div class="actions">
        <span class="delete-button"><i class="fas fa-trash"></i></span>
        <span class="edit-button"><i class="fas fa-edit"></i></span>
    </div>
    `;

    listContainer.appendChild(listOfTask);
    listContainer.appendChild(hr);
    inputTask.value = "";

    const checkbox = listOfTask.querySelector("input");
    const taskSpan = listOfTask.querySelector("label span");
    const editButton = listOfTask.querySelector(".edit-button");
    const deleteButton = listOfTask.querySelector(".delete-button");

    checkbox.addEventListener("click", function () {
        listOfTask.classList.toggle("completed");
        updateTask();
    });

    editButton.addEventListener("click", function () {
        const update = prompt("Edit the following task: ", taskSpan.textContent);
        if (update !== null) {
            taskSpan.textContent = update;
            updateTask();
        }
    });

    deleteButton.addEventListener("click", function () {
        if (confirm("Are you sure you want to delete this task? ")) {
            listOfTask.remove();
            hr.remove();
            updateTask();
        }
    });

    updateTask();
}

submitButton.addEventListener("click", addTask);

inputTask.addEventListener("keypress", function (event) {
    if (event.key === "Enter") {
        addTask();
    }
});

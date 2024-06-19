import React, { ChangeEvent, Component, FormEvent } from "react";
import axios from "axios";
import "semantic-ui-css/semantic.min.css";
import { Card, Header, Form, Input, Icon } from "semantic-ui-react";

// Define the API endpoint where the backend server is running
const endpoint = "http://localhost:9000";

// Define the interface for the component's state
interface ToDoListState {
  task: string;
  items: TaskItem[];
}

// Define interface for each TaskItem
interface TaskItem {
  _id: string;
  task: string;
  status: boolean;
}

class ToDoList extends Component<{}, ToDoListState> {
  constructor(props: {}) {
    super(props);

    // Initialize state with empty task and an empty array of items
    this.state = {
      task: "",
      items: [],
    };
  }

  // Calls getTask function which fetches all task from backend  after component is mounted
  componentDidMount() {
    this.getTask();
  }

  // Fetch all tasks from the backend server
  getTask = () => {
    axios.get(endpoint + "/api/task").then((res) => {
      console.log(res);
      if (res.data) {
        this.setState({
          items: res.data, // Update state with fetched task items
        });
      } else {
        this.setState({
          items: [], // Empty List if no data returned
        });
      }
    });
  };

  // Handler for input change events
  onChange = (event: ChangeEvent<HTMLInputElement>): void => {
    const { name, value } = event.target;
    this.setState((prevState) => ({
      ...prevState,
      [name]: value,
    }));
  };

  // Handler for form submission to create a new task
  onSubmit = (event: FormEvent<HTMLFormElement>): void => {
    event.preventDefault();
    const { task } = this.state;
    if (task != "") {
      console.log(endpoint + "/api/tasks", { task });
      axios
        .post(
          endpoint + "/api/tasks",
          { task },
          {
            headers: {
              "Content-Type": "application/json",
            },
          }
        )
        .then((res) => {
          this.getTask(); // Refresh task list after creating new task
          this.setState({
            task: "", // clears input field to empty
          });
          console.log(res);
        })
        .catch((error) => {
          console.error("Error creating task:", error);
        });
    }
  };

  // Handler to mark a task as undone
  undotask = (id: string) => {
    axios
      .put(
        `${endpoint}/api/undoTask/${id}`,
        {},
        {
          headers: {
            "Content-Type": "application/json",
          },
        }
      )
      .then((res) => {
        console.log(res);
        this.getTask(); // Refresh task list
      })
      .catch((error) => {
        console.error("error undoing task", error);
      });
  };

  // Handler to update a task's status as completed (status : True)
  updateTask = (id: string) => {
    axios
      .put(
        `${endpoint}/api/tasks/${id}`,
        {},
        {
          headers: {
            "Content-Type": "application/json",
          },
        }
      )
      .then((res) => {
        console.log(res);
        this.getTask(); // Refresh the task List
      })
      .catch((error) => {
        console.error("Error updating task:", error);
      });
  };

  // Handler to delete a task
  deleteTask = (id: string) => {
    axios
      .delete(`${endpoint}/api/deleteTask/${id}`, {
        headers: {
          "Content-Type": "application/json",
        },
      })
      .then((res) => {
        console.log(res);
        this.getTask(); // Refreshes the Task list
      })
      .catch((error) => {
        console.error("Error deleting task:", error);
      });
  };

  // Render method to render the component UI
  render(): React.ReactNode {
    return (
      <div>
        <div className="row">
          <Header className="header" color="red">
            TO DO LIST
          </Header>
        </div>
        <div className="row">
          <Form onSubmit={this.onSubmit}>
            <Input
              type="text"
              name="task"
              onChange={this.onChange}
              value={this.state.task}
              fluid
              placeholder="Create task"
            />
            <br />
            <button
              type="submit"
              style={{ backgroundColor: "black", color: "white" }}
            >
              Create task
            </button>
          </Form>
        </div>
        <br />
        <div className="row">
          <Card.Group>
            {this.state.items.map((item) => (
              <Card
                key={item._id}
                color={item.status ? "green" : "yellow"}
                fluid
                className="rough"
              >
                <Card.Content>
                  <Card.Header textAlign="left">
                    <div
                      style={{
                        wordWrap: "break-word",
                        ...(item.status && {
                          textDecorationLine: "line-through",
                        }),
                      }}
                    >
                      {item.task}
                    </div>
                  </Card.Header>
                  <Card.Meta textAlign="right">
                    <Icon
                      name="check circle"
                      color="blue"
                      onClick={() => this.updateTask(item._id)}
                    />
                    <span style={{ paddingRight: 10 }}>Done</span>
                    <Icon
                      name="refresh"
                      color="blue"
                      onClick={() => this.undotask(item._id)}
                    />
                    <span style={{ paddingRight: 10 }}>Undo</span>
                    <Icon
                      name="delete"
                      color="red"
                      onClick={() => this.deleteTask(item._id)}
                    />
                    <span style={{ paddingRight: 10 }}>Delete</span>
                  </Card.Meta>
                </Card.Content>
              </Card>
            ))}
          </Card.Group>
        </div>
      </div>
    );
  }
}

export default ToDoList;

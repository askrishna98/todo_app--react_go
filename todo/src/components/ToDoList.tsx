import React, { ChangeEvent, Component, FormEvent } from "react";
import axios from "axios";
import "semantic-ui-css/semantic.min.css";
import { Card, Header, Form, Input, Icon } from "semantic-ui-react";

const endpoint = "http://localhost:9000";

interface ToDoListState {
  task: string;
  items: TaskItem[];
}

interface TaskItem {
  _id: string;
  task: string;
  status: boolean;
}

class ToDoList extends Component<{}, ToDoListState> {
  constructor(props: {}) {
    super(props);

    this.state = {
      task: "",
      items: [],
    };
  }

  componentDidMount() {
    this.getTask();
  }

  onChange = (event: ChangeEvent<HTMLInputElement>): void => {
    this.setState({
      [event.target.name]: event.target.value,
    } as unknown as Pick<ToDoListState, keyof ToDoListState>);
  };

  getTask = () => {
    axios.get(endpoint + "/api/task").then((res) => {
      console.log(res);
      if (res.data) {
        this.setState({
          items: res.data,
        });
      } else {
        this.setState({
          items: [],
        });
      }
    });
  };

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
          this.getTask();
          this.setState({
            task: "",
          });
          console.log(res);
        })
        .catch((error) => {
          console.error("Error creating task:", error);
        });
    }
  };

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
        this.getTask();
      })
      .catch((error) => {
        console.error("error undoing task", error);
      });
  };

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
        this.getTask();
      })
      .catch((error) => {
        console.error("Error updating task:", error);
      });
  };

  deleteTask = (id: string) => {
    console.log(id);
    axios
      .delete(`${endpoint}/api/deleteTask/${id}`, {
        headers: {
          "Content-Type": "application/json",
        },
      })
      .then((res) => {
        console.log(res);
        this.getTask();
      })
      .catch((error) => {
        console.error("Error deleting task:", error);
      });
  };

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

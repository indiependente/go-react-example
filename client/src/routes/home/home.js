import React from 'react'
class Home extends React.Component {
    constructor(props) {
        super(props)

        this.state = {
            name: '',
            description: ''
        }
    }

    componentDidMount() {
        fetch("/hello")
            .then(resp => {
                return resp.json();
            })
            .then(data => {
                console.log(data);
                this.setState(
                    {
                        ...this.state,
                        name: data.text,
                    }
                );
            })
            .catch(err => {
                console.log(err);
            });
    }

    render() {
        return (
            <div>
                <p style='font-family: Roboto'>
                    Hello {this.state.name}!
                </p>
            </div >
        )
    }
}
export default Home

export default function App(){
    return <h1>App Components</h1>
}

// named export

export function Pst(){
    return <h2>PST Components</h2>
}
export async function getadvice() {
    const res = await fetch('https://api.adviceslip.com/advice')
    const data = await res.json()
    return data.slip.advice


}

function button(props) {
    <button onClick={props.getadvice}></button>
}
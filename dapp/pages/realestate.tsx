import { useEffect, useState } from "react"

const realestate = () => {
    const [acc, setacc] = useState<string>()
    useEffect(() => {
        const ans = localStorage.getItem("account")
        console.log(ans);
        setacc(ans as string)
        const accounts = () => {
            fetch("http://localhost:8000/api/v1/queryAccountList", {
                method: "post",
                body: JSON.stringify({
                    args: [{
                        accountId: ans
                    }]
                }),
                credentials: 'same-origin',
                headers: {
                    "Content-type": `application/x-www-form-urlencoded`,
                },

            })
                .then((resp) => resp.json())
                .then((res: Queryacc) => {
                    console.log(res, "check check");
                })
        }
        accounts()
    }, [])
    return (
        <div>{acc}</div>
    )
}

export default realestate
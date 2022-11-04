import { useEffect, useState } from "react"
import { FiLogOut } from "react-icons/fi"
import { FaUserNinja } from "react-icons/fa"
import { AiOutlinePlusCircle } from "react-icons/ai"
import { useRouter } from "next/router"
import toast, { Toaster } from "react-hot-toast"
const Realestate = () => {
    const router = useRouter()
    const [acc, setacc] = useState<{ accountId: string; balance: number; userName: string }>()
    const [model, setmodel] = useState<boolean>(false)
    const [model2, setmodel2] = useState<boolean>(false)
    const [model3, setmodel3] = useState<boolean>(false)
    const [check, setcheck] = useState<boolean>(true)
    const [error, seterror] = useState<string>("Couldnt not save")
    const [property, setproperty] = useState<string>("")
    const [order, setorder] = useState<"sell" | "donate" | null>(null)
    const [overall, setoverall] = useState<number>(0)
    const [living, setliving] = useState<number>(0)
    const [price, setprice] = useState<number>(0)
    const [validity, setvalidity] = useState<number>(0)
    const [view, setview] = useState<boolean>(false)
    const [view1, setview1] = useState<boolean>(false)
    const [value, setvalue] = useState<any>(null)
    const [value1, setvalue1] = useState<any>(null)
    const [accnt, setaccnt] = useState<{ accountId: string; balance: number; userName: string } | null>(null)
    const [ordacc, setordacc] = useState<{ accountId: string; balance: number; userName: string } | null>(null)
    const [accs, setaccs] = useState<[{ accountId: string; balance: number; userName: string }]>()
    const [RealEstate, setRealEstate] = useState<[QueryRealEstate] | null>(null)



    useEffect(() => {
        const ans = localStorage.getItem("account")
        if (ans == null) {
            router.push("/")
        }
        setacc(JSON.parse(ans as string))
        const accounts = () => {
            fetch("http://localhost:8000/api/v1/queryAccountList", {
                method: "post"
            })
                .then((resp) => resp.json()
                    .then((res: Queryacc) => {
                        console.log(res);
                        setaccs(res.data)
                    })).catch((err) => alert(err))
        }
        accounts()
        QueryRealEstate()
    }, [])


    const logger = () => {
        console.log(RealEstate, "try");
    }

    const SellProperty = () => {
        toast.promise(
            fetch("http://localhost:8000/api/v1/createSelling", {
                method: "post",
                body: JSON.stringify({
                    objectOfSale: property,
                    "price": price,
                    "salePeriod": validity,
                    "seller": acc?.accountId
                })
            }).then((res) => res.json().then((resp) => {
                console.log(resp)
                QueryRealEstate()
                setmodel3(!model3)
            }).catch((err) => seterror(JSON.stringify(err)))),
            {
                loading: 'Saving...',
                success: <b>Settings saved!</b>,
                error: <b>{error}</b>,
            }
        )
    }

    const donateproperty = () => {
        toast.promise(
            fetch("http://localhost:8000/api/v1/createDonating", {
                method: "post",
                body: JSON.stringify({
                    ObjectOfDonating: property,
                    Grantee: ordacc?.accountId,
                    Donor: acc?.accountId
                })
            }).then((res) => res.json().then((resp) => {
                console.log(resp)
                QueryRealEstate()
                setmodel3(!model3)
            }).catch((err) => seterror(JSON.stringify(err)))),
            {
                loading: 'Saving...',
                success: <b>Settings saved!</b>,
                error: <b>{error}</b>,
            }
        )
    }

    const CreateAsset = () => {
        console.log(accnt, living, overall);
        toast.promise(
            fetch("http://localhost:8000/api/v1/createRealEstate", {
                method: "post",
                body: JSON.stringify({
                    accountId: acc?.accountId,
                    proprietor: accnt?.accountId,
                    totalArea: overall,
                    livingSpace: living
                })
            }).then((res) => res.json().then((resp) => {
                console.log(resp)
                QueryRealEstate()
            }).catch((err) => seterror(JSON.stringify(err)))),
            {
                loading: 'Saving...',
                success: <b>Settings saved!</b>,
                error: <b>{error}</b>,
            }
        )
    }

    const QueryRealEstate = () => {
        fetch("http://localhost:8000/api/v1/queryRealEstateList", {
            method: "post",
        }).then((res) => res.json().then((resp) => {
            setRealEstate(resp.data)
        }).catch((err) => console.log(err)))

    }


    return (
        <div className={`flex flex-col dark:bg-gray-900 overflow-x-hidden overflow-y-scroll h-screen w-screen items-center ${check ? "" : "main"}`}>

            <div className="navbar sticky top-0 z-50 text-white flex justify-between px-10 items-center w-full dark:bg-gray-700 dark:border-gray-700 dark:text-white py-4">
                <div className="flex space-x-6">
                    <button className="glass-button" onClick={() => router.push("/selling")}>Sellings</button>
                    <button className="glass-button" onClick={() => router.push("/donation")}>Donations</button>
                </div>
                <div className="font-semibold text-xl">
                    Block Chain Based Real Estate
                </div>
                <div className="flex items-center space-x-11">
                    <input type="checkbox" defaultChecked onClick={() => {
                        setcheck(!check)
                        localStorage.setItem("dark", check.toString())
                    }} className="hidden" id="darkmode-toggle" />
                    <label className={` ${check ? "" : "avoiddark"} swtch w-16 h-7 `} htmlFor="darkmode-toggle" />
                    <div className="font-bold p-5 hover:bg-slate-400 hover:text-black bg-slate-500 rounded-lg cursor-pointer" onClick={() => {
                        localStorage.removeItem("account")
                        router.push("/")
                    }}>
                        <FiLogOut className="text-xl " />
                    </div>
                </div>
            </div>
            <div className="w-full text-slate-200 text-2xl font-thin underline font-mono pt-10 px-8">
                Your Organization Products are listed Here
            </div>
            <div className=" flex relative flex-col text-white  items-center w-screen h-screen">
                <div >
                    {
                        RealEstate == null ? <div>The Listed Properties are Shown here</div> : (
                            <div className="grid grid-cols-3 gap-16 p-12">
                                {
                                    RealEstate.map((ele, index) => <div className={`${acc?.accountId == "5feceb66ffc8" || acc?.accountId == ele.proprietor ? "" : "hidden"}`} key={index}>
                                        <div className={`w-full bg-white rounded-lg shadow dark:border md:mt-0 sm:max-w-md xl:p-0 dark:bg-gray-800 dark:border-gray-700`}>
                                            <div className="p-6 space-y-4 md:space-y-6 sm:p-8">
                                                <h1 className="text-lg italic font-thin leading-tight tracking-tight text-gray-700 md:text-2xl dark:text-white">
                                                    Listed  Property No. {index + 1}
                                                </h1>
                                                <div className="space-y-4 md:space-y-6" >
                                                    <div className="flex justify-center w-full">
                                                        <div className='w-full'>
                                                            <div className="dropdown relative w-full space-y-2">
                                                                <button
                                                                    className=" dropdown-toggle  px-6 py-2.5   font-medium text-xs leading-tight uppercase  shadow-md   focus:shadow-lg focus:outline-none focus:ring-0   active:shadow-lg active:text-white transition duration-150 ease-in-out  items-center whitespace-nowrap  justify-center bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 space-y-7 dark:placeholder-gray-400 dark:text-white "
                                                                    type="button"
                                                                    id="dropdownMenuButton1"
                                                                >
                                                                    <div>Property ID : {ele.realEstateId}</div>
                                                                    <div>Onwer    ID : {ele.proprietor}</div>
                                                                    <div>OverAll Space : {ele.totalArea}</div>
                                                                    <div>Living Space : {ele.livingSpace}</div>
                                                                </button>
                                                                <button
                                                                    className=" dropdown-toggle  px-6 py-2.5   font-medium text-xs leading-tight uppercase  shadow-md   focus:shadow-lg focus:outline-none focus:ring-0   active:shadow-lg active:text-white transition duration-150 ease-in-out  items-center whitespace-nowrap  justify-center bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 space-y-7 dark:placeholder-gray-400 dark:text-white "
                                                                    type="button"
                                                                    id="dropdownMenuButton1"
                                                                >
                                                                    Guarantee  {ele.encumbrance.toString()}
                                                                </button>
                                                                <div className={`${acc?.accountId == "5feceb66ffc8" || ele.encumbrance == true ? "hidden" : ""} flex w-full justify-between px-5 py-4`}>
                                                                    <div className="bg-green-700 cursor-pointer px-7 py-3 rounded-xl" onClick={() => {
                                                                        setorder("sell")
                                                                        setmodel3(!model3)
                                                                        setproperty(ele.realEstateId)
                                                                    }}>Sell</div>
                                                                    <div className="bg-blue-600 cursor-pointer px-7 py-3 rounded-xl" onClick={() => {
                                                                        setorder("donate")
                                                                        setmodel3(!model3)
                                                                        setproperty(ele.realEstateId)
                                                                    }}>Donate</div>
                                                                </div>
                                                            </div>
                                                        </div>
                                                    </div>
                                                </div>
                                            </div>
                                        </div>
                                    </div>)
                                }
                            </div>
                        )
                    }
                </div>
            </div>
            <button
                type='button'
                className='fixed bottom-7 right-11 inline-block  rounded-full bg-green-500 p-3 text-xs font-medium uppercase leading-tight text-white shadow-md transition duration-150 ease-in-out hover:bg-green-600 hover:shadow-lg focus:bg-green-400 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-green-800 active:shadow-lg '
                onClick={() => setmodel(!model)}
            >
                <FaUserNinja className='h-8 w-8' />
            </button>
            <button
                type='button'
                className={`fixed bottom-7 left-11 inline-block  rounded-full bg-blue-500 p-3 text-xs font-medium uppercase leading-tight text-white shadow-md transition duration-150 ease-in-out hover:bg-blue-600 hover:shadow-lg focus:bg-blue-400 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-green-800 active:shadow-lg ${acc?.userName == "administrator" ? "" : "hidden"
                    }`}
                onClick={() => setmodel2(!model2)}
            >
                <AiOutlinePlusCircle className='h-8 w-8' />
            </button>
            <div className={`fixed z-10 overflow-y-auto top-0 w-full left-0 ${model ? "" : "hidden"}`} id="modal">
                <div className="flex items-center justify-center min-height-100vh pt-4 px-4 pb-20 text-center sm:block sm:p-0">
                    <div className="fixed inset-0 transition-opacity">
                        <div className="absolute inset-0 bg-gray-900 opacity-75" />
                    </div>
                    <span className="hidden sm:inline-block sm:align-middle sm:h-screen">&#8203;</span>
                    <div className="inline-block align-center bg-white rounded-lg text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-lg sm:w-full" role="dialog" aria-modal="true" aria-labelledby="modal-headline">
                        <div className="bg-white px-4 text-lg font-semibold text-slate-600 pt-5 pb-4 sm:p-6 sm:pb-4">
                            <label>Account Details</label>
                            <div className="px-3 py-4 bg-gray-200 rounded-xl m-5">
                                <div className="flex">
                                    <div>Account ID :</div>
                                    <div>{acc?.accountId}</div>
                                </div>
                                <div className="flex">
                                    <div>Balence Held :</div>
                                    <div>{acc?.balance}</div>
                                </div>
                                <div className="flex">
                                    <div>UserName :</div>
                                    <div>{acc?.userName}</div>
                                </div>
                            </div>
                            <div>
                                <i>
                                    After the sales, donation or pledge operations, the guarantee status is true

                                    When the guarantee status is false, the sale, donation or pledge operation can be initiated
                                </i>
                            </div>
                        </div>
                        <div className="bg-gray-200 px-4 py-3 w-full flex justify-between">
                            <a href="https://github.com/jayendramadaram/Hyperledger-Based-real-estate">
                                <button type="button" className="py-2 px-4 bg-green-500 text-white rounded hover:bg-gray-700 mr-2"><i className="fas fa-times"></i> VISIT PROJECT</button>
                            </a>
                            <button type="button" className="py-2 px-4 bg-gray-500 text-white rounded hover:bg-gray-700 mr-2" onClick={() => setmodel(!model)} ><i className="fas fa-times"></i> Cancel</button>
                        </div>
                    </div>
                </div>
            </div>
            <div className={`fixed z-10 overflow-y-auto top-0 w-full left-0 ${model2 ? "" : "hidden"}`} id="modal">
                <div className="flex items-center justify-center min-height-100vh pt-4 px-4 pb-20 text-center sm:block sm:p-0">
                    <div className="fixed inset-0 transition-opacity">
                        <div className="absolute inset-0 bg-gray-900 opacity-75" />
                    </div>
                    <span className="hidden sm:inline-block sm:align-middle sm:h-screen">&#8203;</span>
                    <div className="inline-block align-center bg-white rounded-lg text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-lg sm:w-full" role="dialog" aria-modal="true" aria-labelledby="modal-headline">
                        <div className="bg-white px-4 text-lg font-semibold text-slate-600 pt-5 pb-4 sm:p-6 sm:pb-4">
                            <label>Mint New Property to User</label>
                            <div className="px-3 py-4 bg-gray-200 rounded-xl m-5">
                                <div className="space-y-4 md:space-y-6" >
                                    <div className="flex justify-center w-full">
                                        <div className='w-full'>
                                            <div className="dropdown relative w-full">
                                                <button
                                                    className=" dropdown-toggle  px-6 py-2.5 font-medium text-xs leading-tight uppercase  shadow-md  hover:bg-blue-700 hover:shadow-lg  focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0  active:bg-blue-800 active:shadow-lg active:text-white transition duration-150 ease-in-out  items-center whitespace-nowrap  justify-center bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                                                    type="button"
                                                    id="dropdownMenuButton1"
                                                    data-bs-toggle="dropdown"
                                                    aria-expanded="false"
                                                    onClick={() => setview(!view)}
                                                >
                                                    <div className='flex'>
                                                        <div>{accnt == null ? "Selected None" : value}</div>
                                                        <svg
                                                            aria-hidden="true"
                                                            focusable="false"
                                                            data-prefix="fas"
                                                            data-icon="caret-down"
                                                            className="w-2 ml-2"
                                                            role="img"
                                                            xmlns="http://www.w3.org/2000/svg"
                                                            viewBox="0 0 320 512"
                                                        >
                                                            <path
                                                                fill="currentColor"
                                                                d="M31.3 192h257.3c17.8 0 26.7 21.5 14.1 34.1L174.1 354.8c-7.8 7.8-20.5 7.8-28.3 0L17.2 226.1C4.6 213.5 13.5 192 31.3 192z"
                                                            ></path>
                                                        </svg>
                                                    </div>
                                                </button>
                                                <ul
                                                    className={`dropdown-menu min-w-max absolute  bg-white text-base z-50 float-left py-2 list-none text-left rounded-lg shadow-lg mt-1 m-0 bg-clip-padding border-none w-full ${view ? "" : "hidden"}`}
                                                    aria-labelledby="dropdownMenuButton1"
                                                >
                                                    {
                                                        accs?.map((ele, index) => ele.userName == "administrator" ? "" : <li key={index} onClick={() => {
                                                            setvalue(ele.userName)
                                                            setview(!view)
                                                            setaccnt(ele)
                                                        }}>
                                                            <a
                                                                className="dropdown-item text-sm py-2 px-4 font-normal block w-full whitespace-nowrap bg-transparent  text-gray-700  hover:bg-gray-100"
                                                                href="#"
                                                            >{ele.userName}
                                                            </a>
                                                        </li>)
                                                    }


                                                </ul>
                                            </div>
                                        </div>
                                    </div>
                                    <div className="flex justify-between px-3 font-semibold">
                                        <div>OverALL space m^2</div>
                                        <input
                                            type="number"
                                            className=" form-control block w-full px-3 py-1.5 text-base font-normal  text-gray-700  bg-white bg-clip-padding  border border-solid border-gray-300  rounded transition ease-in-out  m-0  focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none"
                                            id="exampleNumber0"
                                            placeholder="Number input"
                                            step={0.1}
                                            min={0}
                                            value={overall}
                                            onChange={(e) => setoverall(Number(e.target.value))}
                                        />
                                    </div>
                                    <div className="flex justify-between px-3 font-semibold">
                                        <div>Living space m^2</div>
                                        <input
                                            type="number"
                                            className=" form-control block w-full px-3 py-1.5 text-base font-normal  text-gray-700  bg-white bg-clip-padding  border border-solid border-gray-300  rounded transition ease-in-out  m-0  focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none"
                                            id="exampleNumber0"
                                            placeholder="Number input"
                                            step={0.1}
                                            min={0}
                                            value={living}
                                            onChange={(e) => setliving(Number(e.target.value))}
                                        />
                                    </div>
                                </div>
                            </div>

                        </div>
                        <div className="bg-gray-200 px-4 py-3 w-full flex justify-between">
                            <button type="button" className="py-2 px-4 bg-green-500 text-white rounded hover:bg-gray-700 mr-2"
                                onClick={() => CreateAsset()}
                            ><i className="fas fa-times"></i> Create </button>

                            <button type="button" className="py-2 px-4 bg-gray-500 text-white rounded hover:bg-gray-700 mr-2" onClick={() => setmodel2(!model2)} ><i className="fas fa-times"></i> Cancel</button>
                        </div>
                    </div>
                </div>
                <div className="z-50">
                    <Toaster
                        position="bottom-left"
                        reverseOrder={false}
                    />
                </div>
            </div>
            <div className={`fixed z-10 overflow-y-auto top-0 w-full  left-0 ${model3 ? "" : "hidden"}`} id="modal">
                <div className="flex items-center justify-center min-height-100vh pt-4 px-4 pb-20 text-center sm:block sm:p-0">
                    <div className="fixed inset-0 transition-opacity">
                        <div className="absolute inset-0 bg-gray-900 opacity-75" />
                    </div>
                    <span className="hidden sm:inline-block sm:align-middle sm:h-screen">&#8203;</span>
                    <div className="inline-block align-center bg-white rounded-lg text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-lg sm:w-full" role="dialog" aria-modal="true" aria-labelledby="modal-headline">
                        <div className={`bg-white px-4 text-lg font-semibold text-slate-600 pt-5 pb-4 sm:p-6 sm:pb-4 ${order == "sell" ? "" : "pt-10"}`}>
                            <label>Execute Order</label>
                            {
                                order == "sell" ? <div className="px-3 py-4 bg-gray-200 space-y-6 rounded-xl m-5">
                                    <div className="flex justify-between  px-3 font-semibold">
                                        <div>Price</div>
                                        <input
                                            type="number"
                                            className=" form-control block  px-3 py-1.5 text-base font-normal  text-gray-700  bg-white bg-clip-padding  border border-solid border-gray-300  rounded transition ease-in-out  m-0  focus:text-gray-700 w-3/4 focus:bg-white focus:border-blue-600 focus:outline-none"
                                            id="exampleNumber0"
                                            placeholder="Number input"
                                            step={1000}
                                            min={0}
                                            value={price}
                                            onChange={(e) => setprice(Number(e.target.value))}
                                        />
                                    </div>
                                    <div className="flex justify-between px-3 font-semibold">
                                        <div>Validity</div>
                                        <input
                                            type="number"
                                            className=" form-control block w-3/4  px-3 py-1.5 text-base font-normal  text-gray-700  bg-white bg-clip-padding  border border-solid border-gray-300  rounded transition ease-in-out  m-0  focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none"
                                            id="exampleNumber0"
                                            placeholder="Number input"
                                            step={1}
                                            min={0}
                                            value={validity}
                                            onChange={(e) => setvalidity(Number(e.target.value))}
                                        />
                                    </div>
                                </div> :
                                    <div className="px-3 py-4 bg-gray-200 rounded-xl m-5 pb-32">
                                        <div className="space-y-4 md:space-y-6" >
                                            <div className="flex justify-center w-full">
                                                <div className='w-full'>
                                                    <div className="dropdown relative w-full">
                                                        <button
                                                            className=" dropdown-toggle  px-6 py-2.5 font-medium text-xs leading-tight uppercase  shadow-md  hover:bg-blue-700 hover:shadow-lg  focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0  active:bg-blue-800 active:shadow-lg active:text-white transition duration-150 ease-in-out  items-center whitespace-nowrap  justify-center bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                                                            type="button"
                                                            id="dropdownMenuButton1"
                                                            data-bs-toggle="dropdown"
                                                            aria-expanded="false"
                                                            onClick={() => setview1(!view1)}
                                                        >
                                                            <div className='flex'>
                                                                <div>{ordacc == null ? "Selected None" : value1}</div>
                                                                <svg
                                                                    aria-hidden="true"
                                                                    focusable="false"
                                                                    data-prefix="fas"
                                                                    data-icon="caret-down"
                                                                    className="w-2 ml-2"
                                                                    role="img"
                                                                    xmlns="http://www.w3.org/2000/svg"
                                                                    viewBox="0 0 320 512"
                                                                >
                                                                    <path
                                                                        fill="currentColor"
                                                                        d="M31.3 192h257.3c17.8 0 26.7 21.5 14.1 34.1L174.1 354.8c-7.8 7.8-20.5 7.8-28.3 0L17.2 226.1C4.6 213.5 13.5 192 31.3 192z"
                                                                    ></path>
                                                                </svg>
                                                            </div>
                                                        </button>
                                                        <ul
                                                            className={`dropdown-menu min-w-max absolute  bg-white text-base z-50 float-left py-2 list-none text-left rounded-lg shadow-lg mt-1 m-0 bg-clip-padding border-none w-full ${view1 ? "" : "hidden"}`}
                                                            aria-labelledby="dropdownMenuButton1"
                                                        >
                                                            {
                                                                accs?.map((ele, index) => ele.userName == "administrator" || ele.accountId == acc?.accountId ? "" : <li key={index} onClick={() => {
                                                                    setvalue1(ele.userName)
                                                                    setview1(!view1)
                                                                    setordacc(ele)
                                                                }}>
                                                                    <a
                                                                        className="dropdown-item text-sm py-2 px-4 font-normal block w-full whitespace-nowrap bg-transparent  text-gray-700  hover:bg-gray-100"
                                                                        href="#"
                                                                    >{ele.userName}
                                                                    </a>
                                                                </li>)
                                                            }


                                                        </ul>
                                                    </div>
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                            }

                        </div>
                        <div className="bg-gray-200 px-4 py-3 w-full flex justify-between">
                            <button type="button"
                                onClick={() => order == "sell" ? SellProperty() : donateproperty()}
                                className="py-2 px-4 bg-green-500 text-white rounded hover:bg-gray-700 mr-2"><i className="fas fa-times"></i>{order == "sell" ? "sell" : "donate"}</button>
                            <button type="button" className="py-2 px-4 bg-gray-500 text-white rounded hover:bg-gray-700 mr-2" onClick={() => setmodel3(!model3)} ><i className="fas fa-times"></i> Cancel</button>
                        </div>
                    </div>
                </div>
            </div>
            <Toaster
                position="bottom-left"
                reverseOrder={false}
            />
        </div >
    )
}

export default Realestate

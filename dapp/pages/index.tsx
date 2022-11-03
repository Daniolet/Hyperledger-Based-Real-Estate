import type { NextPage } from 'next'
import { useRouter } from 'next/router'

import { useState, useEffect } from 'react'

const Home: NextPage = () => {
  const [view, setview] = useState<boolean>(false)
  const [value, setvalue] = useState<any>(null)
  const [accnt, setaccnt] = useState<any>(null)
  const [accs, setaccs] = useState<[{ accountId: string; balance: number; userName: string }]>()
  const router = useRouter()
  useEffect(() => {
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
  }, []);

  return (
    <div>
      <section className="bg-gray-50 dark:bg-gray-900">
        <div className="flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0">
          <a href="#" className="flex items-center mb-6 text-2xl font-semibold text-gray-900 dark:text-white">
            <img className="w-8 h-8 mr-2" src="https://i.pinimg.com/originals/90/25/57/9025576bc65d3395157e05b987eea548.jpg" alt="logo" />
            HyperLedger Real Estate
          </a>
          <div className="w-full bg-white rounded-lg shadow dark:border md:mt-0 sm:max-w-md xl:p-0 dark:bg-gray-800 dark:border-gray-700">
            <div className="p-6 space-y-4 md:space-y-6 sm:p-8">
              <h1 className="text-xl font-bold leading-tight tracking-tight text-gray-900 md:text-2xl dark:text-white">
                Sign in to your account
              </h1>
              <form className="space-y-4 md:space-y-6" action="#">
                <div className="flex justify-center w-full">
                  <div className='w-full'>
                    <div className="dropdown relative w-full">
                      <button
                        className=" dropdown-toggle  px-6 py-2.5   font-medium text-xs leading-tight uppercase  shadow-md  hover:bg-blue-700 hover:shadow-lg  focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0  active:bg-blue-800 active:shadow-lg active:text-white transition duration-150 ease-in-out  items-center whitespace-nowrap  justify-center bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
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
                        className={`dropdown-menu min-w-max absolute bg-white text-base z-50 float-left py-2 list-none text-left rounded-lg shadow-lg mt-1 m-0 bg-clip-padding border-none w-full ${view ? "" : "hidden"}`}
                        aria-labelledby="dropdownMenuButton1"
                      >
                        {
                          accs?.map((ele, index) => <li key={index} onClick={() => {
                            setvalue(ele.userName)
                            setview(!view)
                            setaccnt(ele)
                          }}>
                            <a
                              className="dropdown-item text-sm py-2 px-4 font-normal block w-full whitespace-nowrap bg-transparent  text-gray-700  hover:bg-gray-100
            "
                              href="#"
                            >{ele.userName}
                            </a>
                          </li>)
                        }


                      </ul>
                    </div>
                  </div>
                </div>



                <div className="flex items-center justify-between">
                  <div className="flex items-start">
                    <div className="flex items-center h-5">
                      <input id="remember" type="checkbox" className="w-4 h-4 border border-gray-300 rounded bg-gray-50 focus:ring-3 focus:ring-primary-300 dark:bg-gray-700 dark:border-gray-600 dark:focus:ring-primary-600 dark:ring-offset-gray-800" />
                    </div>
                    <div className="ml-3 text-sm">
                      <label className="text-gray-500 dark:text-gray-300">Remember me</label>
                    </div>
                  </div>
                  {/* <a href="#" className="text-sm font-medium text-primary-600 hover:underline dark:text-primary-500">Forgot password?</a> */}
                </div>
                <button type="submit" className="w-full text-white bg-primary-600 hover:bg-primary-700 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800" onClick={() => {
                  if (value != null) {
                    localStorage.setItem("account", JSON.stringify(accnt))
                    router.push("/realestate")
                    // localStorage.setItem("accOBJ" :)
                  } else alert("Select any Owner")
                }}>Sign in</button>
                <p className="text-sm font-light text-gray-500 dark:text-gray-400">
                  Don’t have an account yet? <a href="mailto:dimebeatengreen8@gmail.com" className="font-medium text-primary-600 hover:underline dark:text-primary-500">Contact Administrator or MSP</a>
                </p>
              </form>
            </div>
          </div>
        </div>
      </section >
    </div >
  )
}

export default Home

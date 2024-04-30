"use client"
import Lottie from "react-lottie";
import animationData from "./Animation - 1713984821166.json";
export function Loader(){
  const defaultOptions = {
    loop: true,
    autoplay: true,
    animationData: animationData,
    rendererSettings: {
      preserveAspectRatio: "xMidYMid slice",
    },
  };
  return(
    <>
          <Lottie options={defaultOptions} height={100} width={100} />
    </>

  )
}

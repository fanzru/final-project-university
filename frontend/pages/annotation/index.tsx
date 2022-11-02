import type { NextPage } from 'next'
import AnnotationComponent from "../../components/annotation"
import Head from 'next/head'
import { useRouter } from 'next/router'

const Login : NextPage = () => {
  const router = useRouter()
  const { paper_id } = router.query
  return (
    <div>
      <AnnotationComponent paper_id={paper_id}/>
    </div>
  )
}

export default Login
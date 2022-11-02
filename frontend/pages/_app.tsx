import '../styles/globals.css'
import '../styles/layer.css'
import type { AppProps } from 'next/app'
import Script from 'next/script';
import { NextSeo } from 'next-seo';
import { MantineProvider } from '@mantine/core';
import { ThemeProvider } from 'next-themes';
import clsx from 'clsx';
import {
  AnimatePresence,
  domAnimation,
  LazyMotion,
  m,
} from 'framer-motion';
import 'react-toastify/dist/ReactToastify.css';
import { ToastContainer } from 'react-toastify';
import Navbar from '../components/navbar/index';
import { Router } from 'next/router';
import NProgress from 'nprogress';
import '../styles/nprogress.css';
import { Worker } from '@react-pdf-viewer/core';

Router.events.on('routeChangeStart', () => NProgress.start());
Router.events.on('routeChangeComplete', () => NProgress.done());
Router.events.on('routeChangeError', () => NProgress.done());

import { variants } from '../animations/variants';

function MyApp({ Component, pageProps,router }: AppProps) {
  return <div>
    <ThemeProvider defaultTheme='white' attribute='class'>
      <NextSeo
        title="Fanzru - Final Project University"
        description="Jangan lupa berusaha walaupun belum tentu berhasil karena menyerah sama saja dengan kalah - Fanzru"
        openGraph={{
          type: 'website',
          url: 'https://skripsi.fanzru.dev/',
          title: 'Fanzru - Final Project University',
          description:
            'Jangan lupa berusaha walaupun belum tentu berhasil karena menyerah sama saja dengan kalah - Fanzru',
          images: [
            {
              url: '/default.jpeg',
            },
          ],
        }}
      />
      <LazyMotion features={domAnimation}>
        <MantineProvider withGlobalStyles withNormalizeCSS>
          <ToastContainer pauseOnFocusLoss={false} />
          <Worker workerUrl='https://unpkg.com/pdfjs-dist@2.14.305/build/pdf.worker.min.js'>
            <div className="min-h-screen flex flex-col h-full">
              <Navbar/>
              <AnimatePresence
                mode='wait'
                initial={false}
                onExitComplete={() => window.scrollTo(0, 0)}
              >
                <m.div
                  key={router.asPath}
                  variants={variants}
                  initial='hidden'
                  animate='enter'
                  exit='exit'
                  transition={{ ease: 'easeInOut', duration: 0.5 }}
                  className={clsx('flex flex-col h-full flex-grow')}
                >
                  <Component {...pageProps} />
                </m.div>
              </AnimatePresence>
            </div>
          </Worker>
        </MantineProvider>
      </LazyMotion>
    </ThemeProvider>
  </div>
}

export default MyApp

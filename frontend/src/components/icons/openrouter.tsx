import type { SVGProps } from 'react';

export default function OpenRouter(props: SVGProps<SVGSVGElement>) {
    return (
        <svg
            fill="none"
            viewBox="0 0 24 24"
            xmlns="http://www.w3.org/2000/svg"
            {...props}
        >
            <path
                d="M2.75 12h9.1c1.2 0 2.17-.97 2.17-2.17V7.7c0-2.64 2.14-4.78 4.78-4.78h2.45"
                stroke="currentColor"
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth="2"
            />
            <path
                d="M2.75 12h9.1c1.2 0 2.17.97 2.17 2.17v2.13c0 2.64 2.14 4.78 4.78 4.78h2.45"
                stroke="currentColor"
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth="2"
            />
            <path
                d="M18.2 6.05 21.25 3 18.2.95M18.2 21.05 21.25 18l-3.05-3.05"
                stroke="currentColor"
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth="2"
            />
        </svg>
    );
}

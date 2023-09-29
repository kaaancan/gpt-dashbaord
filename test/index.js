const { default: OpenAI } = require("openai");


/**
 * This needs to be the proxy host, its the address of the server that is running the go-proxy
 * @type {string}
 */
const proxyHost = process.env.PROXY_HOST || "http://localhost:1007/";

/**
 * This is the version of the openai api that the proxy is using
 */
const OPENAPI_VERSION = "v1";

console.log("Using proxy host: " + proxyHost);

const openai = new OpenAI({
    apiKey: process.env.OPENAI_API_KEY,
    baseURL: proxyHost,
});

async  function main() {
    const resp = await openai.chat.completions.create({
        model: 'gpt-3.5-turbo',
        messages: [{
            role: 'assistant',
            content: 'You are a helpful assistant'
        },
            {
                content: "Hello, how are you?",
                role: 'user'
            },
        ],
    })
    console.log(resp.choices[0].message.content);
}

/**
 * After you run this, you should be able to the log in the dashboard
 */
main()
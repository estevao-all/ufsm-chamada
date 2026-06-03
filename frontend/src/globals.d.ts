declare global {
    interface Window {
        dwr: {
            _: [{
                engine: {
                    remote: {
                        handleCallback: (batchId: string, callId: string, reply: any) => void
                    }
                }
            }]
        }
    }
}

export {}

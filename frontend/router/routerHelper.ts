import {RouteConfig} from 'vue-router'

export default class RouteGroup {
    private readonly parent?: RouteGroup;
    private readonly prefix: string;
    private readonly config: RouteConfig[];

    constructor(prefix: string, config: RouteConfig[], parent?: RouteGroup) {
        this.parent = parent;

        if (prefix == null) {
            prefix = ""
        }

        if (!prefix.startsWith("/")) {
            prefix = "/" + prefix;
        }

        while (prefix.endsWith("/")) {
            prefix = prefix.slice(0, -1)
        }
        this.prefix = prefix;

        const fPath = this.fullPrefix();
        for (const currConf of config) {
            if (!currConf.path.startsWith("/")) {
                currConf.path = "/" + currConf.path;
            }

            currConf.path = fPath + currConf.path;
        }

        this.config = config;
    }

    fullPrefix(): string {
        let ret = "";

        if (this.parent != null) {
            ret = this.parent.fullPrefix()
        }

        return ret + this.prefix;
    }

    getConfig(): RouteConfig[] {
        return this.config;
    }
}
// layouts enum of different layout available
const layouts = {
    DEFAULT: 0
};

export const Layouts = layouts;

/**
 * Return a layout depending of given LayIdx
 * @param {number} layIdx enum val of layouts
 * @return {string} layout name
 */
export function GetLayout(layIdx) {
    switch (layIdx) {
        case layouts.DEFAULT:
            return 'default-layout'
    }
}
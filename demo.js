// Create the first (outer) custom element
class OuterElement extends HTMLElement {
    constructor() {
        super();
        const shadow1 = this.attachShadow({ mode: 'open' });
        
        // Create container for first shadow root
        const div1 = document.createElement('div');
        div1.innerHTML = '<p>First Shadow Root</p>';
        
        // Create and append middle-element to first shadow
        const middleElement = document.createElement('middle-element');
        div1.appendChild(middleElement);
        
        shadow1.appendChild(div1);
    }
}

// Create the second (middle) custom element
class MiddleElement extends HTMLElement {
    constructor() {
        super();
        const shadow2 = this.attachShadow({ mode: 'open' });
        
        // Create container for second shadow root
        const div2 = document.createElement('div');
        div2.innerHTML = '<p>Second Shadow Root</p>';
        
        // Create and append inner-element to second shadow
        const innerElement = document.createElement('inner-element');
        div2.appendChild(innerElement);
        
        shadow2.appendChild(div2);
    }
}

// Create the third (innermost) custom element
class InnerElement extends HTMLElement {
    constructor() {
        super();
        const shadow3 = this.attachShadow({ mode: 'open' });
        
        // Create container for third shadow root
        const div3 = document.createElement('div');
        div3.innerHTML = '<p>Third Shadow Root</p>';
        
        // Add button to demonstrate access to parent shadow roots
        const button = document.createElement('button');
        button.textContent = 'Access Parent Shadow Roots';
        button.addEventListener('click', this.accessParentShadows.bind(this));
        
        div3.appendChild(button);
        shadow3.appendChild(div3);
    }
    
    accessParentShadows() {
        // Method 1: Access parent shadow roots using getRootNode() and host
        const thirdShadow = this.shadowRoot;
        const secondShadow = this.getRootNode().host.getRootNode();
        const firstShadow = secondShadow.host.getRootNode();
        
        console.log('Third Shadow Root:', thirdShadow);
        console.log('Second Shadow Root:', secondShadow);
        console.log('First Shadow Root:', firstShadow);
        
        // Method 2: Direct parent access using closest()
        const parentMiddleElement = this.closest('middle-element');
        const parentOuterElement = this.closest('outer-element');
        
        if (parentMiddleElement) {
            console.log('Parent Middle Element Shadow:', parentMiddleElement.shadowRoot);
        }
        if (parentOuterElement) {
            console.log('Parent Outer Element Shadow:', parentOuterElement.shadowRoot);
        }
    }
}

// Register custom elements
customElements.define('outer-element', OuterElement);
customElements.define('middle-element', MiddleElement);
customElements.define('inner-element', InnerElement);

// Usage example
document.body.innerHTML = '<outer-element></outer-element>';
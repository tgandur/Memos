import convertResourceToDataURL from "./convertResourceToDataURL";

const applyStyles = async (sourceElement: HTMLElement, clonedElement: HTMLElement) => {
  if (!sourceElement || !clonedElement) {
    return;
  }

  if (sourceElement.tagName === "IMG") {
    try {
      const url = await convertResourceToDataURL(sourceElement.getAttribute("src") ?? "");
      (clonedElement as HTMLImageElement).src = url;
    } catch (error) {
      // do nth
    }
  }

  const sourceStyles = window.getComputedStyle(sourceElement);
  for (const item of sourceStyles) {
    clonedElement.style.setProperty(item, sourceStyles.getPropertyValue(item), sourceStyles.getPropertyPriority(item));
  }

  for (let i = 0; i < clonedElement.childElementCount; i++) {
    await applyStyles(sourceElement.children[i] as HTMLElement, clonedElement.children[i] as HTMLElement);
  }
};

const getCloneStyledElement = async (element: HTMLElement) => {
  const clonedElementContainer = document.createElement(element.tagName);
  clonedElementContainer.innerHTML = element.innerHTML;

  await applyStyles(element, clonedElementContainer);

  return clonedElementContainer;
};

export default getCloneStyledElement;

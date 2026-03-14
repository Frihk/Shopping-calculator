const state = {
  items: [],
};

const form = document.getElementById("item-form");
const nameInput = document.getElementById("item-name");
const qtyInput = document.getElementById("item-qty");
const priceInput = document.getElementById("item-price");
const itemsList = document.getElementById("items-list");
const totalQty = document.getElementById("total-qty");
const totalCost = document.getElementById("total-cost");
const statusEl = document.getElementById("status");
const suggestionsList = document.getElementById("suggestions-list");
const checkoutBtn = document.getElementById("checkout");
const clearBtn = document.getElementById("clear-items");

const formatter = new Intl.NumberFormat(undefined, {
  minimumFractionDigits: 2,
  maximumFractionDigits: 2,
});

function setStatus(message, tone) {
  statusEl.textContent = message;
  statusEl.classList.remove("success", "error");
  if (tone) {
    statusEl.classList.add(tone);
  }
}

function formatMoney(value) {
  return formatter.format(value || 0);
}

function renderItems() {
  itemsList.innerHTML = "";

  if (state.items.length === 0) {
    const empty = document.createElement("div");
    empty.className = "list-empty";
    empty.textContent = "No items yet. Add your first item to begin.";
    itemsList.appendChild(empty);
    return;
  }

  state.items.forEach((item, index) => {
    const row = document.createElement("div");
    row.className = "list-row";
    row.innerHTML = `
      <span>${item.name}</span>
      <span>${item.quantity}</span>
      <span>${formatMoney(item.price)}</span>
      <span>${formatMoney(item.quantity * item.price)}</span>
      <button type="button" data-remove="${index}" aria-label="Remove item">×</button>
    `;
    itemsList.appendChild(row);
  });
}

function updateTotals() {
  const totals = state.items.reduce(
    (acc, item) => {
      acc.qty += item.quantity;
      acc.cost += item.quantity * item.price;
      return acc;
    },
    { qty: 0, cost: 0 }
  );

  totalQty.textContent = totals.qty.toLocaleString();
  totalCost.textContent = formatMoney(totals.cost);
}

function renderSuggestions(items) {
  suggestionsList.innerHTML = "";

  if (!items || items.length === 0) {
    const empty = document.createElement("div");
    empty.className = "list-empty";
    empty.textContent = "No suggestions yet.";
    suggestionsList.appendChild(empty);
    return;
  }

  items.forEach((item) => {
    const row = document.createElement("div");
    row.className = "suggestion-item";
    row.innerHTML = `
      <span>${item.name}</span>
      <span>${formatMoney(item.price)} • ${item.freq}x</span>
    `;
    suggestionsList.appendChild(row);
  });
}

function resetForm() {
  nameInput.value = "";
  qtyInput.value = "1";
  priceInput.value = "";
  nameInput.focus();
}

form.addEventListener("submit", (event) => {
  event.preventDefault();
  const name = nameInput.value.trim();
  const quantity = Number.parseInt(qtyInput.value, 10);
  const price = Number.parseFloat(priceInput.value);

  if (!name) {
    setStatus("Item name is required.", "error");
    return;
  }
  if (!Number.isFinite(quantity) || quantity <= 0) {
    setStatus("Quantity must be greater than 0.", "error");
    return;
  }
  if (!Number.isFinite(price) || price <= 0) {
    setStatus("Price must be greater than 0.", "error");
    return;
  }

  state.items.push({ name, quantity, price });
  renderItems();
  updateTotals();
  setStatus("Item added. Ready to finalize.", "success");
  resetForm();
});

itemsList.addEventListener("click", (event) => {
  const target = event.target;
  if (!(target instanceof HTMLElement)) {
    return;
  }
  const index = target.getAttribute("data-remove");
  if (index === null) {
    return;
  }
  const itemIndex = Number.parseInt(index, 10);
  if (Number.isNaN(itemIndex)) {
    return;
  }
  state.items.splice(itemIndex, 1);
  renderItems();
  updateTotals();
  setStatus("Item removed.", "success");
});

clearBtn.addEventListener("click", () => {
  state.items = [];
  renderItems();
  updateTotals();
  setStatus("List cleared.", "success");
});

checkoutBtn.addEventListener("click", async () => {
  if (state.items.length === 0) {
    setStatus("Add at least one item before finalizing.", "error");
    return;
  }

  checkoutBtn.disabled = true;
  setStatus("Saving to backend...", "");

  try {
    const response = await fetch("/api/checkout", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ items: state.items }),
    });

    const data = await response.json();

    if (!response.ok) {
      throw new Error(data.error || "Checkout failed");
    }

    totalQty.textContent = data.totalQuantity.toLocaleString();
    totalCost.textContent = formatMoney(data.totalCost);
    renderSuggestions(data.suggestions || []);
    setStatus("Saved. Suggestions refreshed.", "success");
  } catch (error) {
    setStatus(`Backend error: ${error.message}`, "error");
  } finally {
    checkoutBtn.disabled = false;
  }
});

async function loadSuggestions() {
  try {
    const response = await fetch("/api/suggestions");
    if (!response.ok) {
      throw new Error("Suggestions unavailable");
    }
    const data = await response.json();
    renderSuggestions(data);
    setStatus("Backend status: connected.", "success");
  } catch (error) {
    renderSuggestions([]);
    setStatus("Backend status: not running.", "error");
  }
}

renderItems();
updateTotals();
loadSuggestions();

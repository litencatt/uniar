// 管理画面用JavaScript

document.addEventListener('DOMContentLoaded', function() {
    // コラプス機能の初期化
    initializeCollapse();
});

function initializeCollapse() {
    // data-bs-toggle="collapse"属性を持つすべてのボタンを取得
    const collapseButtons = document.querySelectorAll('[data-bs-toggle="collapse"]');

    collapseButtons.forEach(button => {
        button.addEventListener('click', function(e) {
            e.preventDefault();

            // ターゲット要素を取得
            const targetId = this.getAttribute('data-bs-target');
            const targetElement = document.querySelector(targetId);

            if (targetElement) {
                toggleCollapse(targetElement);

                // aria-expanded属性を更新
                const isExpanded = targetElement.classList.contains('show');
                this.setAttribute('aria-expanded', isExpanded);

                // ボタンのテキストやスタイルを更新（オプション）
                updateButtonAppearance(this, isExpanded);
            }
        });
    });
}

function toggleCollapse(element) {
    if (element.classList.contains('show')) {
        // 閉じる
        element.classList.remove('show');

        // アニメーション効果のために少し遅延
        setTimeout(() => {
            if (!element.classList.contains('show')) {
                element.style.display = 'none';
            }
        }, 300);
    } else {
        // 開く
        element.style.display = 'block';
        // リフローを強制して、アニメーションが確実に動作するようにする
        element.offsetHeight;
        element.classList.add('show');
    }
}

function updateButtonAppearance(button, isExpanded) {
    // ボタンの見た目を更新（オプション）
    if (isExpanded) {
        button.style.transform = 'rotate(0deg)';
    } else {
        button.style.transform = 'rotate(-90deg)';
    }
}

// より汎用的なコラプス関数（他の用途にも使用可能）
window.toggleElement = function(targetSelector) {
    const element = document.querySelector(targetSelector);
    if (element) {
        toggleCollapse(element);
    }
};